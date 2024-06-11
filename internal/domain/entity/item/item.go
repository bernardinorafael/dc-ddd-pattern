package item

import "github.com/google/uuid"

// Item is an entity composing order aggregate
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func New(name string, price float64) (*Item, error) {
	item := Item{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}

	return &item, nil
}
