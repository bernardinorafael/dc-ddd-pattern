package product

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

func New(name string, price float64, quantity int64) (*Product, error) {
	product := Product{
		ID:       uuid.NewString(),
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
	if err := product.validate(); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) validate() error {
	if len(p.Name) == 0 {
		return errors.New("product name cannot be empty")
	}
	if p.Price < 1 {
		return errors.New("price must be greater than zero")
	}
	if p.Quantity < 1 {
		return errors.New("quantity must be greather than zero")
	}
	return nil
}

func (p *Product) ChangePrice(price float64) error {
	p.Price = price

	if err := p.validate(); err != nil {
		return err
	}
	return nil
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
