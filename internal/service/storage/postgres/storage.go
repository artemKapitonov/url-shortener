package postgres

import (
	"github.com/artemKapitonov/url-shortener/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Storage struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool, log *slog.Logger) *Storage {
	return &Storage{db: db, log: log}
}

func (s *Storage) Get(url entity.ShortURL) (entity.FullURL, error) {
	panic("Implement me!!")
}

func (s *Storage) Create(url entity.FullURL) (entity.ShortURL, error) {
	panic("Implement me!!")
}

func (s *Storage) Close() {
	const op = "postgres.Close"

	log := s.log.With(slog.String("op", op))

	s.db.Close()

	log.Info("Postgres database successfully closed")

}
