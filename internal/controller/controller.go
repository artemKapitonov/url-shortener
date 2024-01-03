package controller

import (
	grpcUrlShortner "github.com/artemKapitonov/url-shortner/internal/controller/grpc/v1"
	"github.com/artemKapitonov/url-shortner/internal/service"
	url_shortner_v1 "github.com/artemKapitonov/url-shortner/pkg/url-shortner_v1"
)

type Controller struct {
	grpcUrlShortner.GRPCserverAPI
}

func New(service *service.Service) *Controller {
	return &Controller{
		GRPCserverAPI: *grpcUrlShortner.NewGRPCServerAPI(service, &url_shortner_v1.UnimplementedURLShortnerServer{}),
	}
}
