package app

import (
	"github.com/artemKapitonov/url-shortener/pkg/client/postgresql_client"
	"github.com/artemKapitonov/url-shortener/pkg/client/redis_client"
	"github.com/artemKapitonov/url-shortener/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

func getLoggerConfig() (logging.Config, error) {
	var cfg logging.Config

	err := cleanenv.ReadConfig("config/logger-config.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getServersConfig() (serversCfg, error) {
	var cfg serversCfg

	err := cleanenv.ReadConfig("config/servers-config.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getRedisConfig() (redis_client.Config, error) {
	var cfg redis_client.Config

	err := cleanenv.ReadConfig("config/redis.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getPostgresConfig() (postgresql_client.Config, error) {
	var cfg postgresql_client.Config

	err := cleanenv.ReadConfig("config/postgres.yaml", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
