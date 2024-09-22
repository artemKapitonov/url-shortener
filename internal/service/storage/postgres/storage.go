package postgres

import (
	"context"

	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Get(ctx context.Context, url entity.URL) (entity.URL, error) {
	panic("Implement me!!")
}

func (s *Storage) Create(ctx context.Context, url entity.URL) error {
	panic("Implement me!!")
}

func (s *Storage) Close() error {
	s.db.Close()
	return nil
}
