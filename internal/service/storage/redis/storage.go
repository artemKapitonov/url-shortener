package redis

import (
	"context"
	"fmt"

	"github.com/artemKapitonov/url-shortener/internal/entity"
	goredis "github.com/redis/go-redis/v9"
)

type Storage struct {
	db *goredis.Client
}

func NewStorage(db *goredis.Client) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Get(ctx context.Context, url entity.URL) (entity.URL, error) {
	var err error

	url.FullURL, err = s.db.Get(ctx, url.ShortURL).Result()
	if err != nil {
		return entity.URL{}, fmt.Errorf("Failed to get full url %s", err.Error())
	}

	return url, nil
}

func (s *Storage) Create(ctx context.Context, url entity.URL) error {
	if err := s.db.Set(ctx, url.ShortURL, url.FullURL, 0).Err(); err != nil {
		return fmt.Errorf("Failed to set url %s", err.Error())
	}

	return nil
}

func (s *Storage) Close() error {

	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("Failed to close redis database: %s", err.Error())
	}

	return nil
}
