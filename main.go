package main

import (
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
)

func main() {
	_, err := customer.New("rafael bernardino", "rafael@gmail.com")
	if err != nil {
		panic(err)
	}
}
