package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sethvargo/go-limiter"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/uatu"
	"github.com/uatu/config"
	redisstore "github.com/uatu/internal/storage/redis"
	_ "github.com/uatu/swagger"
	"go.uber.org/zap"
)

type Server struct {
	cfg        config.Config
	logger     *zap.Logger
	quotes     uatu.QuoteRepository
	chains     uatu.ChainRepository
	priceCache *redisstore.RedisService
	rateLimit  func(http.Handler) http.Handler
	rateStore  limiter.Store
}

func New(
	cfg config.Config,
	logger *zap.Logger,
	quotes uatu.QuoteRepository,
	chains uatu.ChainRepository,
	priceCache *redisstore.RedisService,
) (*Server, error) {
	tokens, interval, err := rateLimitOptions(cfg.RateLimit)
	if err != nil {
		return nil, err
	}
	rateStore, err := redisstore.NewRateLimitStore(cfg.Redis, tokens, interval)
	if err != nil {
		return nil, err
	}
	rateLimit, err := newRateLimiter(cfg.RateLimit, rateStore)
	if err != nil {
		_ = rateStore.Close(context.Background())
		return nil, err
	}

	return &Server{
		cfg:        cfg,
		logger:     logger,
		quotes:     quotes,
		chains:     chains,
		priceCache: priceCache,
		rateLimit:  rateLimit,
		rateStore:  rateStore,
	}, nil
}

func (s *Server) Run() error {
	defer func() {
		if err := s.rateStore.Close(context.Background()); err != nil {
			s.logger.Warn("Failed to close rate limiter store", zap.Error(err))
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: s.cfg.AllowedOrigins,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		ExposedHeaders: []string{
			"X-Request-ID", "X-Trace-ID",
			"X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset", "Retry-After",
		},
		AllowCredentials: false,
		MaxAge:           86400,
	}))

	r.Use(s.rateLimit)

	r.Route("/quotes", s.quoteRoutes)
	r.Route("/blockchains", s.chainRoutes)

	// Serves the UI at /swagger/index.html and the spec at /swagger/doc.json.
	// The wildcard does not match a bare /swagger, so redirect that explicitly
	// rather than letting the obvious URL 404.
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	addr := ":" + s.cfg.PORT
	s.logger.Info("Listening on port", zap.String("port", addr))
	return http.ListenAndServe(addr, r)
}

func (s *Server) quoteRoutes(r chi.Router) {
	quote := &quoteHandler{
		cfg:        s.cfg,
		quoteRepo:  s.quotes,
		chainRepo:  s.chains,
		priceCache: s.priceCache,
	}
	r.Post("/", WrapHTTPHandler(s.logger, quote.CreateQuote, s.cfg, "CreateQuote"))
	r.Post("/routes", WrapHTTPHandler(s.logger, quote.GetQuotes, s.cfg, "GetQuotes"))
}

func (s *Server) chainRoutes(r chi.Router) {
	chain := &chainHandler{
		chainRepo: s.chains,
	}
	r.Get("/", WrapHTTPHandler(s.logger, chain.GetBlockchains, s.cfg, "GetBlockchains"))
	r.Get("/tokens", WrapHTTPHandler(s.logger, chain.GetTokens, s.cfg, "GetTokens"))
	r.Get("/pools", WrapHTTPHandler(s.logger, chain.GetPools, s.cfg, "GetPools"))
	r.Get("/dex", WrapHTTPHandler(s.logger, chain.GetDex, s.cfg, "GetDex"))
}
