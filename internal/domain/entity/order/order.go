package order

import (
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/google/uuid"
)

type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Items      []item.Item `json:"items"`
}

func New(customerID string, items []item.Item) (*Order, error) {
	order := Order{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		Items:      items,
	}

	return &order, nil
}
