package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	*URLShortnerPsql
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{URLShortnerPsql: NewURLShortnerPsql(db)}
}
