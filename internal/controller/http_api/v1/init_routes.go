package http_api

import (
	"log/slog"

	"github.com/artemKapitonov/url-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type logFormat struct {
	Time   string `json:"time,omitempty"`
	Method string `json:"method,omitempty"`
	Uri    string `json:"uri,omitempty"`
	Status string `json:"status,omitempty"`
	Err    string `json:"err,omitempty"`
}

type HttpServerApi struct {
	log *slog.Logger
	UrlService
}

func New(s *service.Service, log *slog.Logger) *HttpServerApi {
	return &HttpServerApi{
		log:        log,
		UrlService: s,
	}
}

func (api *HttpServerApi) InitRoutes(logger *slog.Logger) *echo.Echo {
	var router = echo.New()

	router.Use(Logger(logger))

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
