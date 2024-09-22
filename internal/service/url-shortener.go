package service

import (
	"context"
	"fmt"
	"time"

	"github.com/artemKapitonov/url-shortener/internal/entity"
)

type UrlShortener interface {
	Get(ctx context.Context, url entity.URL) (entity.URL, error)
	Create(ctx context.Context, url entity.URL) error
}

func (s *Service) Get(ctx context.Context, url entity.URL) (entity.URL, error) {
	url.ShortURL = "http://localhost:8080/" + url.ShortURL
	return s.UrlShortener.Get(ctx, url)
}

func (s *Service) Create(ctx context.Context, url entity.URL) (entity.URL, error) {
	url.ShortURL = generateShortURL()
	if err := s.UrlShortener.Create(ctx, url); err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func generateShortURL( /*url string*/ ) string {
	//TODO: make algorithm for short urls
	return fmt.Sprintf("http://localhost:8080/%d", time.Now().UnixNano())
}
