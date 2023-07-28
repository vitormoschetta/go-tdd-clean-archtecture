package mock

import (
	"errors"
	"go-tdd-clean/11/domain/product"
	"time"
)

type ProductRepositoryFake struct {
	storage []product.Product
	err     error
}

func NewProductRepositoryFake() *ProductRepositoryFake {
	return &ProductRepositoryFake{
		storage: []product.Product{},
	}
}

func (r *ProductRepositoryFake) Save(p product.Product) error {
	if r.err != nil {
		return r.err
	}
	r.storage = append(r.storage, p)
	return nil
}

func (r *ProductRepositoryFake) Query(fn func(product.Product) bool) (products []product.Product, err error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, item := range r.storage {
		if fn(item) {
			products = append(products, item)
		}
	}
	return
}

func (r *ProductRepositoryFake) SetError() {
	r.err = errors.New("error")
}

func (r *ProductRepositoryFake) Seed() {
	r.storage = []product.Product{
		{
			ID:        "1",
			Name:      "Product 1",
			Price:     100,
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "2",
			Name:      "Product 2",
			Price:     200,
			CreatedAt: string(time.Now().Format("2006-01-02")),
		},
		{
			ID:        "3",
			Name:      "Product 3",
			Price:     300,
			CreatedAt: string(time.Now().Add(-3600 * time.Hour).Format("2006-01-02")),
		},
	}
}
