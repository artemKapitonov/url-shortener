package storage

import (
	"context"

	"github.com/artemKapitonov/url-shortner/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLShortnerPsql struct {
	db *pgxpool.Pool
}

func NewURLShortnerPsql(db *pgxpool.Pool) *URLShortnerPsql {
	return &URLShortnerPsql{db: db}
}

func (us *URLShortnerPsql) Get(context.Context, *entity.ShortURL) (*entity.FullURL, error) {
	panic("Implement me!!")
}

func (us *URLShortnerPsql) Create(context.Context, *entity.FullURL) (*entity.ShortURL, error) {
	panic("Implement me!!")
}
