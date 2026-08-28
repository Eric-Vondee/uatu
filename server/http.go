package server

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/uatu/config"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type HTTPHandler func(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error)

func WrapHTTPHandler(
	logger *zap.Logger,
	handler HTTPHandler,
	cfg config.Config,
	spanName string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span, requestID := getTracer(r.Context(), r, spanName, cfg.Otel.IsEnabled)
		defer span.End()

		if requestID != "" {
			w.Header().Set("X-Request-ID", requestID)
		}
		if spanContext := span.SpanContext(); spanContext.IsValid() {
			w.Header().Set("X-Trace-ID", spanContext.TraceID().String())
		}

		log := logger.With(zap.String("request_id", requestID))

		resp, handlerErr := handler(ctx, span, log, w, r)
		if resp == nil {
			if handlerErr != nil {
				log.Error("Request handler failed", zap.Error(handlerErr))
			}
			resp = APIError{newAPIResponse(
				http.StatusInternalServerError,
				"an internal server error occurred",
				nil,
			)}
		} else if handlerErr != nil {
			log.Debug("Request rejected", zap.Error(handlerErr))
		}

		err := render.Render(w, r, resp)
		if err != nil {
			log.Error("Failed to render response", zap.Error(err))
		}
	}
}
