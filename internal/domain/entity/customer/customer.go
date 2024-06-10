package customer

import (
	"errors"

	"github.com/google/uuid"
)

type Customer struct {
	id      string
	name    string
	address string
	enabled bool
}

func New(name, address string) (*Customer, error) {
	c := Customer{
		id:      uuid.New().String(),
		name:    name,
		enabled: false,
		address: address,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Customer) validate() error {
	if len(c.name) == 0 {
		return errors.New("customer name cannot be empty")
	}

	return nil
}

func (c *Customer) Enable() error {
	if len(c.address) == 0 {
		return errors.New("cannot enable a customer if they doesnt have an address")
	}

	c.enabled = true
	return nil
}

func (c *Customer) Disable() {
	c.enabled = false
}

func (c *Customer) ChangeName(name string) error {
	c.name = name

	if err := c.validate(); err != nil {
		return err
	}

	return nil
}
