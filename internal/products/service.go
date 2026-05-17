package products

import (
	"context"
	"log/slog"
)

type Service interface {
	ListProducts(ctx context.Context, remote_ip string) ([]Products, error)
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context, remote_ip string) ([]Products, error) {
	pr := NewProductRepo()
	products, err := pr.GetProducts(ctx)
	if err != nil {
		slog.Error("gagal GetProducts", "error", err.Error())
	}
	return products[:], nil
}
