package main

import (
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/valueobj/address"
)

func main() {
	addr := address.Address{Street: "street xpto", City: "china", Zip: "889899"}
	_, err := customer.New("rafael bernardino", addr)
	if err != nil {
		panic(err)
	}
}
