package config

import (
	"log/slog"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	LocalPort int    `env:"LOCAL_PORT"`
	DBURL     string `env:"DATABASE_URL"`

	// If "local", runs API on local port
	RunMode string `env:"RUN_MODE"`
}

func LoadConfig() (Environment, error) {
	var cfg Environment

	// load .env
	err := godotenv.Load()
	if err != nil {
		slog.Info("Error loading .env file")
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		slog.Error("problem reading environment config", "err", err)
		return cfg, err
	}

	// Set RunMode
	if cfg.RunMode == "" {
		slog.Error("RUN_MODE not set, defaulting to local")
		cfg.RunMode = "local"
	}

	// Set LocalPort
	if cfg.LocalPort == 0 {
		slog.Error("LOCAL_PORT not set, defaulting to 8080")

		cfg.LocalPort = 8080
		return cfg, nil
	}

	return cfg, nil
}
