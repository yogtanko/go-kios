package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	go func() {
		if err := api.run(api.mount()); err != nil {
			slog.Error("server has failed to start", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server gracefully...")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Closing database connection...")
	api.db.Close()

	slog.Info("Server exited gracefully")
}
