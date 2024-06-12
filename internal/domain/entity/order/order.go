package order

import (
	"errors"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/google/uuid"
)

type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Total      float64     `json:"total"`
	Items      []item.Item `json:"items"`
}

func New(customerID string, items []item.Item) (*Order, error) {
	total := 0.0
	for _, item := range items {
		total += item.Price
	}

	order := Order{
		ID:         uuid.NewString(),
		CustomerID: customerID,
		Items:      items,
		Total:      total,
	}

	if err := order.validate(); err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *Order) validate() error {
	if len(o.CustomerID) == 0 {
		return errors.New("customer ID cannot be empty")
	}

	if len(o.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	return nil
}

func (o *Order) GetTotalItems() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price
	}
	return total
}
