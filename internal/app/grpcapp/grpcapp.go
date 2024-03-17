package grpcapp

import (
	"fmt"
	"github.com/artemKapitonov/url-shortener/internal/controller"
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/v1"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type GRPCApp struct {
	log    *slog.Logger
	port   string
	Server *grpc.Server
}

func NewGRPCServer(controller *controller.Controller, port string, log *slog.Logger) *GRPCApp {
	grpcServer := grpc.NewServer()

	grpc_api.Register(grpcServer, controller.GrpcServerApi)

	return &GRPCApp{
		log:    log,
		port:   port,
		Server: grpcServer,
	}
}

func (a *GRPCApp) RunGRPCServer() error {
	const op = "grpcapp.RunGRPCServer:"

	log := a.log.With(slog.String("op", op))

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
