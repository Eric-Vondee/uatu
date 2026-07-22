package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/uatu"
	"github.com/uatu/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg    config.Config
	logger *zap.Logger
	quotes uatu.QuoteRepository
	chains uatu.ChainRepository
}

func New(
	cfg config.Config,
	logger *zap.Logger,
	quotes uatu.QuoteRepository,
	chains uatu.ChainRepository,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		quotes: quotes,
		chains: chains,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/quotes", s.quoteRoutes)
	r.Route("/blockchains", s.chainRoutes)

	addr := ":" + s.cfg.PORT
	s.logger.Info("Listening on port", zap.String("port", addr))
	return http.ListenAndServe(addr, r)
}

func (s *Server) quoteRoutes(r chi.Router) {
	quote := &quoteHandler{
		cfg:       s.cfg,
		quoteRepo: s.quotes,
		chainRepo: s.chains,
	}
	r.Post("/", WrapHTTPHandler(s.logger, quote.CreateQuote, s.cfg, "CreateQuote"))
}

func (s *Server) chainRoutes(r chi.Router) {
	chain := &chainHandler{
		chainRepo: s.chains,
	}
	r.Get("/", WrapHTTPHandler(s.logger, chain.GetBlockchains, s.cfg, "GetBlockchains"))
	r.Get("/tokens", WrapHTTPHandler(s.logger, chain.GetTokens, s.cfg, "GetTokens"))
	r.Get("/pools", WrapHTTPHandler(s.logger, chain.GetPools, s.cfg, "GetPools"))
}
