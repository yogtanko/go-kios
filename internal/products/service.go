package products

import (
	"context"
	"log/slog"

	"github.com/yogtanko/go-kios/internal/postgress"
)

type Service interface {
	ListProducts(ctx context.Context) ([]Products, error)
}

type svc struct {
	db *postgress.Database
}

func NewService(db *postgress.Database) Service {
	return &svc{
		db: db,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]Products, error) {
	pr := NewProductRepo(s.db)
	products, err := pr.GetProducts(ctx)
	if err != nil {
		slog.Error("gagal GetProducts", "error", err.Error())
	}
	return products[:], nil
}
