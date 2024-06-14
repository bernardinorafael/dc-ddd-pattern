package customer

import (
	"errors"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/vobject/email"
	"github.com/google/uuid"
)

type Customer struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Email   email.Email `json:"email"`
	Rewards int         `json:"rewards"`
	Enabled bool        `json:"enabled"`
}

func New(customerName, customerEmail string) (*Customer, error) {
	customer := Customer{
		ID:      uuid.NewString(),
		Name:    customerName,
		Email:   email.Email(customerEmail),
		Rewards: 0,
		Enabled: false,
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
