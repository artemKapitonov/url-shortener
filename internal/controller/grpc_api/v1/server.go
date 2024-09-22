package grpc_api

import (
	"context"

	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortener/internal/service"

	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
	"google.golang.org/grpc"
)

type GrpcServerApi struct {
	ctx context.Context
	Convertor
	*url_shortener_v1.UnimplementedURLShortenerServer
	UrlService
}

func NewGRPCServerAPI(
	ctx context.Context,
	c *convertor.EntityConvertor,
	unImplServ *url_shortener_v1.UnimplementedURLShortenerServer,
	us *service.Service,
) *GrpcServerApi {
	return &GrpcServerApi{
		ctx:                             ctx,
		Convertor:                       c,
		UnimplementedURLShortenerServer: unImplServ,
		UrlService:                      us,
	}
}

func Register(gRPC *grpc.Server, api *GrpcServerApi) {

	url_shortener_v1.RegisterURLShortenerServer(gRPC, api)
}
