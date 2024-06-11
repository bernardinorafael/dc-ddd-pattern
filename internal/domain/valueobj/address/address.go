package address

import (
	"errors"
	"fmt"
)

type Address struct {
	Street string
	City   string
	Zip    string
}

func New(street, city, zip string) (*Address, error) {
	address := Address{
		Street: street,
		City:   city,
		Zip:    zip,
	}

	if err := address.validate(); err != nil {
		return nil, err
	}

	return &address, nil
}

func (a *Address) validate() error {
	if len(a.Street) == 0 {
		return errors.New("street cannot be empty")
	}
	if len(a.City) == 0 {
		return errors.New("city cannot be empty")
	}
	if len(a.Zip) == 0 {
		return errors.New("zip cannot be empty")
	}

	return nil
}

func (a *Address) ToString() string {
	return fmt.Sprintf("%s, %s - %s", a.Street, a.City, a.Zip)
}
