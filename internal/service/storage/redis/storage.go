package redis

import (
	"github.com/artemKapitonov/url-shortener/internal/entity"
	goredis "github.com/redis/go-redis/v9"
	"log/slog"
)

type Storage struct {
	log *slog.Logger
	db  *goredis.Client
}

func NewStorage(db *goredis.Client, log *slog.Logger) *Storage {
	return &Storage{db: db, log: log}
}

func (s *Storage) Get(url entity.ShortURL) (entity.FullURL, error) {
	panic("Implement me!!")
}

func (s *Storage) Create(url entity.FullURL) (entity.ShortURL, error) {
	panic("Implement me!!")
}

func (s *Storage) Close() {
	const op = "redis.Close"

	log := s.log.With(slog.String("op", op))

	err := s.db.Close()
	if err != nil {
		log.Error("Failed to close redis database Error:", err)
	}

	log.Info("Redis database successfully closed")
}
