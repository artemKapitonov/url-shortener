package grpc_api

import (
	"github.com/artemKapitonov/url-shortner/internal/service"
	url_shortner_v1 "github.com/artemKapitonov/url-shortner/pkg/url-shortner_v1"
	"google.golang.org/grpc"
)

type GrpcServerApi struct {
	convertor Convertor
	*url_shortner_v1.UnimplementedURLShortnerServer
	service *service.Service
}

type Convertor interface {
	Convert()
}

type GrpcOptions struct {
	Service    *service.Service
	UnImplServ *url_shortner_v1.UnimplementedURLShortnerServer
	Convertor
}

func NewGRPCServerAPI(optins GrpcOptions) *GrpcServerApi {

	return &GrpcServerApi{
		convertor:                      optins.Convertor,
		UnimplementedURLShortnerServer: optins.UnImplServ,
		service:                        optins.Service,
	}
}

func Register(gRPC *grpc.Server) {
	url_shortner_v1.RegisterURLShortnerServer(gRPC, &GrpcServerApi{})
}
