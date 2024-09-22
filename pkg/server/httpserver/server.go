package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/artemKapitonov/url-shortener/pkg/logging"
	"golang.org/x/sync/errgroup"
)

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
)

// Server with HTTP protocol.
type Server struct {
	ctx    context.Context
	server *http.Server
}

// New is creating new http server.
func New(ctx context.Context, handler http.Handler, port string) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         ":" + port,
	}

	s := &Server{
		ctx:    ctx,
		server: httpServer,
	}

	return s
}

// Start is starting http server.
func (s *Server) Start() error {
	log := logging.LoggerFromContext(s.ctx)

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
