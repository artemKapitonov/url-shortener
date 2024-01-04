package controller

import (
	"github.com/artemKapitonov/url-shortner/internal/controller/grpc_api"
	"github.com/artemKapitonov/url-shortner/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortner/internal/controller/http_api/v1"
	"github.com/artemKapitonov/url-shortner/internal/service"
	url_shortner_v1 "github.com/artemKapitonov/url-shortner/pkg/url-shortner_v1"
)

type Controller struct {
	grpc_api.GrpcServerApi
	http_api.HttpServerApi
}

func New(s *service.Service) *Controller {
	return &Controller{
		GrpcServerApi: *grpc_api.NewGRPCServerAPI(grpc_api.GrpcOptions{
			Service:    s,
			Convertor:  convertor.New(),
			UnImplServ: &url_shortner_v1.UnimplementedURLShortnerServer{},
		}),
	}
}
