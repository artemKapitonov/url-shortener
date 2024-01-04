package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/artemKapitonov/url-shortner/internal/app/grpcapp"
	"github.com/artemKapitonov/url-shortner/internal/controller"
	"github.com/artemKapitonov/url-shortner/internal/service"
	"github.com/artemKapitonov/url-shortner/internal/service/storage"
	"github.com/artemKapitonov/url-shortner/pkg/client/postgresql"
	"github.com/artemKapitonov/url-shortner/pkg/logging"
	"github.com/artemKapitonov/url-shortner/pkg/server/httpserver"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/sync/errgroup"
)

// App struct of url-shortner application.
type App struct {
	Controller *controller.Controller
	Service    *service.Service
	Storage    *storage.Storage
	HttpServer *httpserver.Server
	GrpcServer *grpcapp.GRPCApp
}

type serversCfg struct {
	GrpsPort string `yaml:"grpc-port"`
	HttpPort string `yaml:"http-port"`
}

// New create new App struct
func New() *App {
	var app App

	var logger = logging.New()

	slog.SetDefault(logger.Logger)

	ctx := context.TODO()

	DBcfg, err := getDBConfig()
	if err != nil {
		slog.Error("Can't get db configs Error:", err)
	}

	ServersCfg, err := getServersConfig()
	if err != nil {
		slog.Error("Can't get servers configs Error:", err)
	}

	db, err := postgresql.ConnectToDB(ctx, DBcfg)
	if err != nil {
		slog.Error("Can't connect to postgres database Error:", err)
	}

	app.Storage = storage.New(db)

	app.Service = service.New(app.Storage)

	app.Controller = controller.New(app.Service)

	app.GrpcServer = grpcapp.NewGRPCServer(ServersCfg.GrpsPort)

	app.HttpServer = httpserver.New(app.Controller.InitRoutes(logger), ServersCfg.HttpPort)

	return &app
}

// Run method of App for runing application.
func (a *App) Run() {
	g := new(errgroup.Group)
	g.Go(a.GrpcServer.RunGRPCServer)
	g.Go(a.HttpServer.Start)

	if err := g.Wait(); err != nil {
		panic(err)
	}

	ShutdownApp(a)
}

func ShutdownApp(a *App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err := a.HttpServer.Shutdown(context.Background())
	if err != nil {
		return err
	}

	a.GrpcServer.Server.GracefulStop()
	if err != nil {
		return err
	}

	return nil
}

func getServersConfig() (serversCfg, error) {
	var cfg serversCfg

	err := cleanenv.ReadConfig("config/servers-config.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getDBConfig() (postgresql.Config, error) {
	var cfg postgresql.Config

	err := cleanenv.ReadConfig("config/db-config.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
