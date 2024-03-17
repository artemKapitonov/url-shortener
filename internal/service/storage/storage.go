package storage

import (
	"errors"
	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/artemKapitonov/url-shortener/internal/service/storage/postgres"
	"github.com/artemKapitonov/url-shortener/internal/service/storage/redis"
	goredis "github.com/redis/go-redis/v9"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	postgresDatabase = "postgres"
	redisDatabase    = "redis"
)

type Client interface {
	Get(url entity.ShortURL) (entity.FullURL, error)
	Create(url entity.FullURL) (entity.ShortURL, error)
	Close()
}

type Storage struct {
	Client
	log *slog.Logger
}

// New ...
func New(logger *slog.Logger, pgPool *pgxpool.Pool, rdb *goredis.Client, dbType string) *Storage {
	const op = "storage.New"
	var storageClient Client

	log := logger.With(slog.String("op", op))

	switch dbType {
	case postgresDatabase:
		storageClient = postgres.NewStorage(pgPool, logger)
	case redisDatabase:
		storageClient = redis.NewStorage(rdb, logger)
	default:
		panic(errors.New("invalid storage flag"))
	}

	return &Storage{
		Client: storageClient,
		log:    log,
	}
}
