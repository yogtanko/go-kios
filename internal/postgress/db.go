package postgress

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(dbUrl *string) (*Database, error) {
	slog.Info(*dbUrl)
	poolConfig, err := pgxpool.ParseConfig(*dbUrl)
	if err != nil {
		slog.Error("Gagal load config DB", "error", err.Error())
		return nil, err
	}
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	poolConfig.MaxConns = 25
	poolConfig.MinConns = 10
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		slog.Error("Gagal membuat pool", "error", err.Error())
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		slog.Error("Gagal terhubung ke server", "error", err.Error())
		return nil, err
	}
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		slog.Info("Database connection pool berhasil di tutup")
	}
}
