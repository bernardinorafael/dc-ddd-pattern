package email

import (
	"errors"
	"strings"
)

type Email string

type email struct {
	value string
}

func New(address string) (*email, error) {
	e := email{value: address}

	if err := e.validate(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *email) validate() error {
	if !strings.Contains(e.value, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func (e *email) GetEmail() Email {
	return Email(e.value)
}
