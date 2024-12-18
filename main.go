package main

import (
	"chomp-cheap-api/config"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	// Setup Logging
	slog.SetDefault(config.SetupLogger(os.Stdout))

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "err", err)
		os.Exit(1)
	}

	// Create server
	s := config.CreateServer(cfg)

	if cfg.RunMode == "local" {
		slog.Info("Starting local execution")
		err = s.Start(":" + strconv.Itoa(cfg.LocalPort))
		if err != nil {
			slog.Error("Failed to start server", "err", err)
			os.Exit(1)
		}
	} else {
		slog.Info("unknown RUNMODE", "mode", cfg.RunMode)
		os.Exit(1)
	}
}
