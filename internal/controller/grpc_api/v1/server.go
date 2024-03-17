package grpc_api

import (
	"log/slog"

	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortener/internal/service"

	url_shortener_v1 "github.com/artemKapitonov/url-shortener/pkg/url-shortener_v1"
	"google.golang.org/grpc"
)

type GrpcServerApi struct {
	log *slog.Logger
	Convertor
	*url_shortener_v1.UnimplementedURLShortenerServer
	UrlService
}

func NewGRPCServerAPI(
	log *slog.Logger,
	c *convertor.EntityConvertor,
	unImplServ *url_shortener_v1.UnimplementedURLShortenerServer,
	us *service.Service,
) *GrpcServerApi {
	return &GrpcServerApi{
		log:                             log,
		Convertor:                       c,
		UnimplementedURLShortenerServer: unImplServ,
		UrlService:                      us,
	}
}

func Register(gRPC *grpc.Server, api *GrpcServerApi) {

	url_shortener_v1.RegisterURLShortenerServer(gRPC, api)
}
