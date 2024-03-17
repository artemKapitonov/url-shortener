package grpc_api

import (
	"context"
	"github.com/artemKapitonov/url-shortener/internal/entity"
	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
)

type UrlService interface {
	Get(url entity.ShortURL) (entity.FullURL, error)
	Create(url entity.FullURL) (entity.ShortURL, error)
}

type Convertor interface {
	Convert(url *url_shortener_v1.ShortURL) entity.ShortURL
}

func (api *GrpcServerApi) Get(ctx context.Context, inputUrl *url_shortener_v1.ShortURL) (*url_shortener_v1.FullURL, error) {
	const op = "grpc_api.GrpcServerApi.Get"

	url := api.Convertor.Convert(inputUrl)

	result, err := api.UrlService.Get(url)
	if err != nil {
		return nil, err
	}

	return &url_shortener_v1.FullURL{Url: result.Url}, nil
}

func (api *GrpcServerApi) Create(context.Context, *url_shortener_v1.FullURL) (*url_shortener_v1.ShortURL, error) {
	panic("implement me")
}
