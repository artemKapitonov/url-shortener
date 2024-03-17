package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
)

// Server with HTTP protocol.
type Server struct {
	log    *slog.Logger
	server *http.Server
}

// New is creating new http server.
func New(handler http.Handler, port string, log *slog.Logger) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         ":" + port,
	}

	s := &Server{
		log:    log,
		server: httpServer,
	}

	return s
}

// Start is starting http server.
func (s *Server) Start() error {
	const op = "httpserver.Start:"

	log := s.log.With(slog.String("op", op))

	var err error

	log.Info(fmt.Sprintf("HTTP server started at address %s", s.server.Addr))

	g := new(errgroup.Group)

	g.Go(s.server.ListenAndServe)

	return err
}

// Shutdown id stopping http server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
