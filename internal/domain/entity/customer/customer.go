package customer

import (
	"errors"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/valueobj/address"
	"github.com/google/uuid"
)

type Customer struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Address address.Address `json:"address"`
	Rewards int             `json:"rewards"`
	Enabled bool            `json:"enabled"`
}

func New(name string, addr address.Address) (*Customer, error) {
	customer := Customer{
		ID:      uuid.New().String(),
		Name:    name,
		Enabled: false,
		Address: addr,
		Rewards: 0,
	}
	if err := customer.validate(); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *Customer) validate() error {
	if len(c.Name) == 0 {
		return errors.New("customer name cannot be empty")
	}
	return nil
}

func (c *Customer) Enable() {
	c.Enabled = true
}

func (c *Customer) Disable() {
	c.Enabled = false
}

func (c *Customer) ChangeName(name string) error {
	c.Name = name
	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Customer) IncreaseRewards(rewards int) error {
	if rewards < 1 {
		return errors.New("rewards points must be greater than zero")
	}
	c.Rewards += rewards

	return nil
}
