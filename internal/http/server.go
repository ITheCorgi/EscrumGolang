package server

import (
	"context"
	"net/http"

	"github.com/ITheCorgi/EscrumGolang/internal/config"
)

type Server struct {
	httpSrv *http.Server
}

//New returns a new HTTP Server Instance with pre-configured data
func New(cfg *config.ConfigMap, handler http.Handler) *Server {
	return &Server{
		httpSrv: &http.Server{
			Addr:           ":" + cfg.HTTPData.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTPData.ReadTimeout,
			WriteTimeout:   cfg.HTTPData.WriteTimeout,
			MaxHeaderBytes: cfg.HTTPData.MaxHeaderBytes,
		},
	}
}

//Start function calls built-in for running server on specified path
func (s *Server) Start() error {
	return s.httpSrv.ListenAndServe()
}

//Shutdown closes http server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
