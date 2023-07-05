package mock

import (
	"errors"
	"go-tdd-clean/10/domain/product"
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

func (r *ProductRepositoryFake) Query(from string, to string) (products []product.Product, err error) {
	if r.err != nil {
		return nil, r.err
	}
	for _, item := range r.storage {
		if item.CreatedAt >= from && item.CreatedAt <= to {
			products = append(products, item)
		}
	}
	return
}

func (r *ProductRepositoryFake) SetError() {
	r.err = errors.New("error")
}
