package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/uatu/config"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const serviceName = "uatu.server"

var tracer = otel.Tracer(serviceName)
var noopTracer = noop.NewTracerProvider().Tracer(serviceName)

type OTelShutdown func(context.Context) error

func getRequestID(r *http.Request) string {
	return middleware.GetReqID(r.Context())
}

func getTracer(
	ctx context.Context,
	r *http.Request,
	operationName string,
	isEnabled bool,
) (context.Context, trace.Span, string) {
	rid := getRequestID(r)
	ctx, span := tracer.Start(ctx, operationName)
	if !isEnabled {
		ctx, span = noopTracer.Start(ctx, operationName)
		return ctx, span, rid
	}
	span.SetAttributes(attribute.String("request_id", rid))
	return ctx, span, rid
}

func initializeResources() (*resource.Resource, error) {
	serviceResources := resource.NewWithAttributes(
		"",
		attribute.String("service.name", serviceName),
		attribute.String("library.language", "go"),
	)
	return resource.Merge(
		resource.Default(),
		serviceResources,
	)
}

func InitOTELCapabilities(
	ctx context.Context,
	cfg config.Config,
) (OTelShutdown, error) {
	if !cfg.Otel.IsEnabled {
		return func(context.Context) error {
			return nil
		}, nil
	}

	endpoint, err := normalizeOTLPEndpoint(
		cfg.Otel.Endpoint,
		cfg.Otel.UseTLS,
	)
	if err != nil {
		return nil, fmt.Errorf("normalize OTLP endpoint: %w", err)
	}
	headers := parseOTLPHeaders(cfg.Otel.Headers)

	res, err := initializeResources()
	if err != nil {
		return nil, fmt.Errorf("initialize OTEL resource: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpointURL(endpoint),
		otlptracegrpc.WithHeaders(headers),
	)
	if err != nil {
		return nil, fmt.Errorf("initialize OTLP trace exporter: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExporter),
	)
	metricExporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpointURL(endpoint),
		otlpmetricgrpc.WithHeaders(headers),
	)
	if err != nil {
		_ = tracerProvider.Shutdown(context.Background())

		return nil, fmt.Errorf("initialize OTLP metric exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(
			metric.NewPeriodicReader(metricExporter),
		),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	if err := runtime.Start(
		runtime.WithMeterProvider(meterProvider),
	); err != nil {
		_ = meterProvider.Shutdown(context.Background())
		_ = tracerProvider.Shutdown(context.Background())

		return nil, fmt.Errorf("start Go runtime metrics: %w", err)
	}

	return newOTelShutdown(
		tracerProvider,
		meterProvider,
	), nil
}

func normalizeOTLPEndpoint(
	endpoint string,
	useTLS bool,
) (string, error) {
	endpoint = strings.TrimSpace(endpoint)
	if endpoint == "" {
		return "", errors.New("OTLP endpoint is empty")
	}

	if !strings.Contains(endpoint, "://") {
		scheme := "https"
		if !useTLS {
			scheme = "http"
		}

		endpoint = scheme + "://" + endpoint
	}

	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("parse endpoint: %w", err)
	}

	switch parsed.Scheme {
	case "http", "https":
	default:
		return "", fmt.Errorf(
			"unsupported endpoint scheme %q",
			parsed.Scheme,
		)
	}

	if parsed.Host == "" {
		return "", errors.New("OTLP endpoint has no host")
	}

	if parsed.RawQuery != "" {
		return "", errors.New("OTLP endpoint must not have a query")
	}

	if parsed.Fragment != "" {
		return "", errors.New("OTLP endpoint must not have a fragment")
	}

	parsed.Path = strings.TrimSuffix(parsed.Path, "/")

	return parsed.String(), nil
}

func parseOTLPHeaders(value string) map[string]string {
	headers := make(map[string]string)

	for _, pair := range strings.Split(value, ",") {
		key, headerValue, found := strings.Cut(pair, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		headerValue = strings.TrimSpace(headerValue)

		if key == "" {
			continue
		}

		headers[key] = headerValue
	}

	return headers
}

func newOTelShutdown(
	tracerProvider *sdktrace.TracerProvider,
	meterProvider *metric.MeterProvider,
) OTelShutdown {
	var (
		once        sync.Once
		shutdownErr error
	)

	return func(ctx context.Context) error {
		once.Do(func() {
			shutdownErr = errors.Join(
				meterProvider.Shutdown(ctx),
				tracerProvider.Shutdown(ctx),
			)
		})

		return shutdownErr
	}
}
