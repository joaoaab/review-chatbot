package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Config struct {
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASS"`
		Name     string `env:"DB_NAME"`
	}

	Server struct {
		Port string `env:"SERVER_PORT"`
	}
}

func NewConfig() *Config {
	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	return &c
}
