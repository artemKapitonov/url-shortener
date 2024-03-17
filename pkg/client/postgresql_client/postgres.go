package postgresql_client

import (
	"context"
	"fmt"
	"time"

	"github.com/artemKapitonov/url-shortener/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config represents the configuration for connecting to Postgres database.
type Config struct {
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"user"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	SSLMode  string `yaml:"sslmode" env:"SSLMode"`
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
