package service

import (
	"log/slog"

	"github.com/artemKapitonov/url-shortener/internal/service/storage"
)

type Service struct {
	log *slog.Logger
	UrlShortener
}

func New(storage *storage.Storage, log *slog.Logger) *Service {
	return &Service{
		UrlShortener: storage.Client,
		log:          log,
	}
}
