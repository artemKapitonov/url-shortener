package http_api

import (
	"context"

	"github.com/artemKapitonov/url-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type HttpServerApi struct {
	ctx context.Context
	UrlService
}

func New(ctx context.Context, s *service.Service) *HttpServerApi {
	return &HttpServerApi{
		ctx:        ctx,
		UrlService: s,
	}
}

func (api *HttpServerApi) InitRoutes() *echo.Echo {
	var router = echo.New()

	router.Use(Logger(api.ctx))

	router.GET("/:url", api.getFullUrl)
	api.initUrlGroup(router)

	return router
}

func (api *HttpServerApi) initUrlGroup(r *echo.Echo) {
	url := r.Group("/url")
	url.POST("", api.createShortUrl)

}
