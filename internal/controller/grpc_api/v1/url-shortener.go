package grpc_api

import (
	"context"
	"log/slog"

	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/artemKapitonov/url-shortener/pkg/logging"
	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
)

type UrlService interface {
	Get(ctx context.Context, url entity.URL) (entity.URL, error)
	Create(ctx context.Context, url entity.URL) (entity.URL, error)
}

type Convertor interface {
	Convert(url *url_shortener_v1.ShortURL) entity.URL
}

func (api *GrpcServerApi) Get(ctx context.Context, inputUrl *url_shortener_v1.ShortURL) (*url_shortener_v1.FullURL, error) {
	url := api.Convertor.Convert(inputUrl)

	log := logging.LoggerFromContext(ctx).With(slog.String("short_url", inputUrl.Url))
	ctx = logging.ContextWithLogger(ctx, log)

	log.Info("Get short url")
	result, err := api.UrlService.Get(ctx, url)
	if err != nil {
		log.Error("Failed to return full url", err)
		return nil, err
	}

	return &url_shortener_v1.FullURL{Url: result.FullURL}, nil
}

func (api *GrpcServerApi) Create(context.Context, *url_shortener_v1.FullURL) (*url_shortener_v1.ShortURL, error) {
	panic("implement me")
}
