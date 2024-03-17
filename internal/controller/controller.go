package controller

import (
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/v1"
	"github.com/artemKapitonov/url-shortener/internal/controller/http_api/v1"
	"github.com/artemKapitonov/url-shortener/internal/service"
	"log/slog"

	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
)

type Controller struct {
	log *slog.Logger
	*grpc_api.GrpcServerApi
	*http_api.HttpServerApi
}

func New(s *service.Service, log *slog.Logger, c *convertor.EntityConvertor) *Controller {
	return &Controller{
		GrpcServerApi: grpc_api.NewGRPCServerAPI(
			log,
			c,
			&url_shortener_v1.UnimplementedURLShortenerServer{},
			s,
		),

		HttpServerApi: http_api.New(s, log),
	}
}
