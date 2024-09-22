package service

import (
	"github.com/artemKapitonov/url-shortener/internal/service/storage"
)

type Service struct {
	UrlShortener
}

func New(storage *storage.Storage) *Service {
	return &Service{

		UrlShortener: storage.Client,
	}
}
