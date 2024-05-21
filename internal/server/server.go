package server

import (
	"context"
	"net/http"
	"server-techno-flow/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(config *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + config.Server.Port,
			Handler:        handler,
			ReadTimeout:    config.Server.ReadTimeout,
			WriteTimeout:   config.Server.WriteTimeout,
			MaxHeaderBytes: config.Server.MaxHeaderMegabytes << 20, // 1 MB

		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
