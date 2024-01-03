package service

import "github.com/artemKapitonov/url-shortner/internal/service/storage"

type Service struct {
	UrlShortner
}

func New(storage *storage.Storage) *Service {
	return &Service{
		UrlShortner: storage.URLShortnerPsql,
	}
}
