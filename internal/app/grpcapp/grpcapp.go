package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	grpcUrlShortner "github.com/artemKapitonov/url-shortner/internal/controller/grpc/v1"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	port   string
	Server *grpc.Server
}

func NewGRPCServer(port string) *GRPCApp {
	grpcServer := grpc.NewServer()

	grpcUrlShortner.Register(grpcServer)

	return &GRPCApp{
		port:   port,
		Server: grpcServer,
	}
}

func (a *GRPCApp) RunGRPCServer() error {

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("gPRC server start on address: %s", a.port))

	err = a.Server.Serve(l)
	if err != nil {
		return err
	}
	return nil
}
