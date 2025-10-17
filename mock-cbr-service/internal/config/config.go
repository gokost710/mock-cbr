package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"ENV"`
	HTTPAddr   string `env:"HTTP_ADDR"`
	GRPCAddr   string `env:"GRPC_ADDR"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func Load() *Config {
	_ = godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return &cfg
}
