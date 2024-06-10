package main

import (
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
)

func main() {
	c, _ := customer.New("rafael bernardino", "street xpto, 12")
	c.ChangeName("wenderson malheiros")
}
