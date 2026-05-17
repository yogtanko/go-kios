package products

import "context"

type ProductRepositories interface {
	GetProducts(ctx context.Context) ([]Products, error)
}

type pr struct {
}

func NewProductRepo() ProductRepositories {
	return &pr{}
}

func (pr *pr) GetProducts(ctx context.Context) ([]Products, error) {
	products := [1]Products{{Id: "1", Name: "Nama 1"}}
	return products[:], nil
}
