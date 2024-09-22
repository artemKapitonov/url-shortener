package grpcapp

import (
	"context"
	"fmt"
	"net"

	"github.com/artemKapitonov/url-shortener/internal/controller"
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/v1"
	"github.com/artemKapitonov/url-shortener/pkg/logging"

	"google.golang.org/grpc"
)

type GRPCApp struct {
	ctx    context.Context
	port   string
	Server *grpc.Server
}

func NewGRPCServer(ctx context.Context, controller *controller.Controller, port string) *GRPCApp {
	grpcServer := grpc.NewServer()

	grpc_api.Register(grpcServer, controller.GrpcServerApi)

	return &GRPCApp{
		ctx:    ctx,
		port:   port,
		Server: grpcServer,
	}
}

func (a *GRPCApp) RunGRPCServer() error {
	log := logging.LoggerFromContext(a.ctx)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("gPRC server start on address :%s", a.port))

	err = a.Server.Serve(l)
	if err != nil {
		return err
	}
	return nil
}
