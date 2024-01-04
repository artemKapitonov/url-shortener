package http_api

import (
	"github.com/artemKapitonov/url-shortner/internal/service"
	"github.com/artemKapitonov/url-shortner/pkg/logging"
	"github.com/labstack/echo/v4"
)

type HttpServerApi struct {
	service service.Service
}

func (api *HttpServerApi) InitRoutes(logger *logging.Logger) *echo.Echo {
	var router = echo.New()

	router.Logger.SetOutput(logger.Writer)

	api.initUrlGroup(router)

	return router
}

func (api *HttpServerApi) initUrlGroup(r *echo.Echo) {
	url := r.Group("/url")
	{
		url.GET("", api.getFullUrl)
		url.POST("", api.createShortUrl)
	}
}
