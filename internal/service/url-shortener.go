package service

import (
	"github.com/artemKapitonov/url-shortener/internal/entity"
)

type UrlShortener interface {
	Get(url entity.ShortURL) (entity.FullURL, error)
	Create(url entity.FullURL) (entity.ShortURL, error)
}

func (s *Service) Get(url entity.ShortURL) (entity.FullURL, error) {
	return entity.FullURL{Url: "Hello world"}, nil
}

func (s *Service) Create(url entity.FullURL) (entity.ShortURL, error) {
	panic("implement me")
}
