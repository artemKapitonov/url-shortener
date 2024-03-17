package http_api

import (
	"log/slog"
	"net/http"

	"github.com/artemKapitonov/url-shortener/internal/entity"

	"github.com/labstack/echo/v4"
)

type UrlService interface {
	Get(url entity.ShortURL) (entity.FullURL, error)
	Create(url entity.FullURL) (entity.ShortURL, error)
}

func (api *HttpServerApi) getFullUrl(ctx echo.Context) error {
	var url entity.ShortURL
	if err := ctx.Bind(&url); err != nil {
		return err
	}

	const op = "http_api.getFullUrl"

	log := api.log.With(
		slog.String("op", op),
		slog.String("url", url.Url),
	)

	log.Info("Get short url")

	fullUrl, err := api.UrlService.Get(url)
	if err != nil {
		return err
	}

	if err := ctx.JSON(http.StatusOK, fullUrl); err != nil {
		return err
	}

	return nil
}

func (api *HttpServerApi) createShortUrl(ctx echo.Context) error {
	panic("implement me")
}
