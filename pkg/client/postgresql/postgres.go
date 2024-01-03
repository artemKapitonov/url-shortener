package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/artemKapitonov/url-shortner/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Client is an interface for PostgresSQL client operations.
type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Ping(ctx context.Context) error
	Close()
}

// Config represents the configuration for connecting to Postgres database.
type Config struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

// connectionConfig creates a connection configuration for Postgres database.
func getConnectionConfig(cfg Config) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ConnectToDB connects to the Postgres database.
func ConnectToDB(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	var db *pgxpool.Pool

	var err error

	var maxAttempts = 5

	connCfg, err := getConnectionConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = utils.DoWithTries(
		func() error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			db, err = pgxpool.NewWithConfig(ctx, connCfg)
			if err != nil {
				return err
			}

			err := db.Ping(ctx)
			if err != nil {
				return err
			}

			return nil
		}, maxAttempts, 5*time.Second)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
