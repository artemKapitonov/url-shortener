package http_api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/artemKapitonov/url-shortener/internal/controller/dto"
	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/artemKapitonov/url-shortener/pkg/logging"

	"github.com/labstack/echo/v4"
)

type UrlService interface {
	Get(ctx context.Context, url entity.URL) (entity.URL, error)
	Create(ctx context.Context, url entity.URL) (entity.URL, error)
}

func (api *HttpServerApi) getFullUrl(ctx echo.Context) error {
	var url dto.ShortURL
	url.Url = ctx.Param("url")

	log := logging.LoggerFromContext(api.ctx).
		With(
			slog.String("url", url.Url),
		)

	log.Info("Get short url")

	fullUrl, err := api.UrlService.Get(api.ctx, url.ConvertToEntity())
	if err != nil {
		return dto.SendError(ctx, err.Error(), http.StatusInternalServerError)
	}

	return ctx.Redirect(http.StatusFound, fullUrl.FullURL)
}

func (api *HttpServerApi) createShortUrl(ctx echo.Context) error {
	var inputURL dto.FullURL
	if err := ctx.Bind(&inputURL); err != nil {
		return err
	}

	log := logging.LoggerFromContext(api.ctx).With(slog.String("short_url", inputURL.Url))
	api.ctx = logging.ContextWithLogger(api.ctx, log)

	log.Info("Get short url")

	fullUrl, err := api.UrlService.Create(api.ctx, inputURL.ConvertToEntity())
	if err != nil {
		return dto.SendError(ctx, err.Error(), http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, fullUrl)
}
