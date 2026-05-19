package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/yogtanko/go-kios/pkg/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Gagal load .env", "error", err)
	}
	cfg := config.GetConfig()

	api := application{
		config: cfg,
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}
	defer func() {
		slog.Info("Closing database connection...")
		api.db.Close()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		slog.Info("Shutting down server gracefully...")
		slog.Info("Server exited gracefully")
	}
}
