package app

import (
	"context"
	"fmt"
	"github.com/artemKapitonov/url-shortener/pkg/client/redis_client"
	"github.com/jackc/pgx/v5/pgxpool"
	goredis "github.com/redis/go-redis/v9"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/artemKapitonov/url-shortener/internal/app/grpcapp"
	"github.com/artemKapitonov/url-shortener/internal/controller"
	"github.com/artemKapitonov/url-shortener/internal/controller/grpc_api/convertor"
	"github.com/artemKapitonov/url-shortener/internal/service"
	"github.com/artemKapitonov/url-shortener/internal/service/storage"
	"github.com/artemKapitonov/url-shortener/pkg/client/postgresql_client"
	"github.com/artemKapitonov/url-shortener/pkg/logging"
	"github.com/artemKapitonov/url-shortener/pkg/server/httpserver"

	"golang.org/x/sync/errgroup"
)

// App struct of url-shortener application.
type App struct {
	log        *slog.Logger
	Controller *controller.Controller
	Service    *service.Service
	Storage    *storage.Storage
	HttpServer *httpserver.Server
	GrpcServer *grpcapp.GRPCApp
}

type serversCfg struct {
	GrpcPort string `yaml:"grpc-port"`
	HttpPort string `yaml:"http-port"`
}

// New create new App struct
func New() *App {
	const op = "app.New:"

	var app App

	ctx := context.TODO()

	loggerCfg, err := getLoggerConfig()

	var logger = logging.New(loggerCfg)

	app.log = logger.Logger

	log := app.log.With(slog.String("op", op))

	postgresCfg, err := getPostgresConfig()
	if err != nil {
		log.Error("Failed to get pool configs Error:", err)
	}

	redisCfg, err := getRedisConfig()
	if err != nil {
		log.Error("Failed to get redis config Error:", err)
	}

	ServersCfg, err := getServersConfig()
	if err != nil {
		log.Error("Failed to get servers configs Error:", err)
	}

	dbType := getTypeOfStorageByFlag()

	var redisDB *goredis.Client
	var pgPool *pgxpool.Pool

	switch dbType {
	case "redis":
		redisDB, err = redis_client.ConnectToDB(ctx, redisCfg)
		if err != nil {
			log.Error("Failed to connect to redis database Error:", err)
		}
	case "postgres":
		pgPool, err = postgresql_client.ConnectToDB(ctx, postgresCfg)
		if err != nil {
			log.Error("Failed to connect to postgres database Error:", err)
			panic(err)
		}
	}

	log.Info(fmt.Sprintf("Successfully connected to the %s database", dbType))

	app.Storage = storage.New(app.log, pgPool, redisDB, dbType)

	app.Service = service.New(app.Storage, app.log)

	conv := convertor.New()

	app.Controller = controller.New(app.Service, app.log, conv)

	app.GrpcServer = grpcapp.NewGRPCServer(app.Controller, ServersCfg.GrpcPort, app.log)

	app.HttpServer = httpserver.New(app.Controller.InitRoutes(logger), ServersCfg.HttpPort, app.log)

	return &app
}

// Run method of App for running application.
func (a *App) Run() error {
	g := new(errgroup.Group)
	g.Go(a.GrpcServer.RunGRPCServer)
	g.Go(a.HttpServer.Start)
	ShutdownApp(a)

	if err := g.Wait(); err != nil {
		return err
	}

	return nil

}

func ShutdownApp(a *App) {
	const op = "app.ShutdownApp:"

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log := a.log.With(slog.String("op", op))

	err := a.HttpServer.Shutdown(context.Background())
	if err != nil {
		log.Warn("HTTP server stop with Error:", err)
	}

	a.GrpcServer.Server.GracefulStop()
	if err != nil {
		log.Warn("gRPC server stop with Error:", err)
	}

	a.Storage.Client.Close()

	log.Info("Application stopped successfully")
}
