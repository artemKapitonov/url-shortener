package logging

import (
	"context"
	"errors"
	"log/slog"
)

// https://dev.to/ilyakaznacheev/where-to-place-logger-in-golang-13o3

type ctxLogger struct{}

// ContextWithLogger adds logger to context.
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// loggerFromContext returns logger from context.
func loggerFromContext(ctx context.Context) (*Logger, error) {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l, nil
	}

	return nil, errors.New("Logger not founded")
}
