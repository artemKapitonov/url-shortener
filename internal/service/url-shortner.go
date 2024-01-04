package service

import (
	"context"

	"github.com/artemKapitonov/url-shortner/internal/entity"
)

type UrlShortner interface {
	Get(context.Context, *entity.ShortURL) (*entity.FullURL, error)
	Create(context.Context, *entity.FullURL) (*entity.ShortURL, error)
}

func Get(context.Context, *entity.ShortURL) (*entity.FullURL, error) {
	panic("implement me")
}

func Create(context.Context, *entity.FullURL) (*entity.ShortURL, error) {
	panic("implement me")
}
