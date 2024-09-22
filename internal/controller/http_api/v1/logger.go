package http_api

import (
	"context"
	"log/slog"
	"time"

	"github.com/artemKapitonov/url-shortener/pkg/logging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger is echo.MiddlewareFunc logger with slog.Logger.
func Logger(ctx context.Context) echo.MiddlewareFunc {
	log := logging.LoggerFromContext(ctx)

	log = log.With(
		slog.String("component", "middleware/logger"),
	)

	log.Info("logger middleware enabled")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			entry := log.With(
				slog.String("method", req.Method),
				slog.String("path", req.URL.Path),
				slog.String("remote_addr", req.RemoteAddr),
				slog.String("user_agent", req.UserAgent()),
				slog.String("request_id", middleware.DefaultRequestIDConfig.Generator()),
			)

			t1 := time.Now()
			defer func() {
				entry.Info("request completed",
					slog.Int("status", res.Status),
					slog.Int("bytes", int(res.Size)),
					slog.String("duration", time.Since(t1).String()),
				)
			}()

			err = next(c)

			return err
		}
	}
}
