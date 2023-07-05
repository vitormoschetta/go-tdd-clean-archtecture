package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        string
	Name      string
	Price     float64
	CreatedAt string
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().Format(time.RFC3339), // 2021-01-01 00:00:00
	}
}

func (p *Product) Validate() (errs []string) {
	if p.Name == "" {
		errs = append(errs, "name is required")
	}
	if p.Price <= 0 {
		errs = append(errs, "price must be greater than 0")
	}
	if p.ID == "" {
		errs = append(errs, "id is required")
	}
	if p.CreatedAt == "" {
		errs = append(errs, "created_at is required")
	}
	return errs
}
