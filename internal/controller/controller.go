package controller

import (
	"context"

	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/v1"
	"github.com/artemKapitonov/url-shortener/internal/controller/http_api/v1"
	"github.com/artemKapitonov/url-shortener/internal/service"

	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
)

type Controller struct {
	ctx context.Context
	*grpc_api.GrpcServerApi
	*http_api.HttpServerApi
}

func New(ctx context.Context, s *service.Service, c *convertor.EntityConvertor) *Controller {

	return &Controller{
		GrpcServerApi: grpc_api.NewGRPCServerAPI(
			ctx,
			c,
			&url_shortener_v1.UnimplementedURLShortenerServer{},
			s,
		),

		HttpServerApi: http_api.New(ctx, s),
	}
}
