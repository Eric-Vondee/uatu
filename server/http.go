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

		log := logger.With(zap.String("request_id", requestID))

		resp, _ := handler(ctx, span, log, w, r)

		err := render.Render(w, r, resp)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
