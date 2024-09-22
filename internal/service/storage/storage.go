package storage

import (
	"context"
	"errors"

	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/artemKapitonov/url-shortener/internal/service/storage/postgres"
	"github.com/artemKapitonov/url-shortener/internal/service/storage/redis"
	goredis "github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	postgresDatabase = "postgres"
	redisDatabase    = "redis"
)

type Client interface {
	Get(ctx context.Context, url entity.URL) (entity.URL, error)
	Create(ctx context.Context, url entity.URL) error
	Close() error
}

type Storage struct {
	Client
}

// New ...
func New(pgPool *pgxpool.Pool, rdb *goredis.Client, dbType string) *Storage {
	const op = "storage.New"
	var storageClient Client

	//log := logging.LoggerFromContext(ctx).With(slog.String("op", op))

	switch dbType {
	case postgresDatabase:
		storageClient = postgres.NewStorage(pgPool)
	case redisDatabase:
		storageClient = redis.NewStorage(rdb)
	default:
		panic(errors.New("invalid storage flag"))
	}

	return &Storage{
		Client: storageClient,
	}
}
