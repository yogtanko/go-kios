package products

import (
	"context"
	"log/slog"

	"github.com/yogtanko/go-kios/internal/postgress"
)

type ProductRepositories interface {
	GetProducts(ctx context.Context) ([]Products, error)
}

type pr struct {
	db postgress.Database
}

func NewProductRepo(db *postgress.Database) ProductRepositories {
	return &pr{
		db: *db,
	}
}

func (pr *pr) GetProducts(ctx context.Context) ([]Products, error) {
	query := `select
  id
  ,code
  ,name
  ,barcode
  ,created_at
from
  "Products"`
	slog.Info("Query: " + query)
	rows, err := pr.db.Pool.Query(ctx, query)
	if err != nil {
		slog.Error("Gagal menjalankan query", "error", err.Error())
	}
	defer rows.Close()

	var products []Products
	i := 1
	for rows.Next() {
		var p Products
		err := rows.Scan(&p.Id, &p.Code, &p.Name, &p.Barcode, &p.CreatedAt)
		if err != nil {
			slog.Error("Gagal membaca row ke " + string(i))
		}
		i++
		products = append(products, p)
	}
	return products, nil
}
