package grpcUrlShortner

import (
	"context"

	"github.com/artemKapitonov/url-shortner/internal/service"
	url_shortner_v1 "github.com/artemKapitonov/url-shortner/pkg/url-shortner_v1"
	"google.golang.org/grpc"
)

type GRPCserverAPI struct {
	*url_shortner_v1.UnimplementedURLShortnerServer
	service *service.Service
}

func NewGRPCServerAPI(s *service.Service, un *url_shortner_v1.UnimplementedURLShortnerServer) *GRPCserverAPI {
	return &GRPCserverAPI{
		UnimplementedURLShortnerServer: un,
		service:                        s,
	}
}

func Register(gRPC *grpc.Server) {
	url_shortner_v1.RegisterURLShortnerServer(gRPC, &GRPCserverAPI{})
}

func (c *GRPCserverAPI) Get(context.Context, *url_shortner_v1.ShortURL) (*url_shortner_v1.FullURL, error) {
	return &url_shortner_v1.FullURL{Url: "eeeee"}, nil
}

func (c *GRPCserverAPI) Create(context.Context, *url_shortner_v1.FullURL) (*url_shortner_v1.ShortURL, error) {
	panic("implement me")
}
