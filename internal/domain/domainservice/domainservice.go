package domainservice

import (
	"errors"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/order"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/product"
)

func IncProductPercent(products []*product.Product, percent int) error {
	if percent > 100 {
		return errors.New("percent value must be between 1 and 100")
	}

	if len(products) == 0 {
		return errors.New("products slice cannot be empty")
	}

	for _, product := range products {
		inc := (product.Price * float64(percent)) / 100
		_ = product.ChangePrice(product.Price + inc)
	}
	return nil
}

func GetOrdersAmount(orders []order.Order) (float64, error) {
	amount := 0.0

	if len(orders) == 0 {
		return amount, errors.New("orders slice cannot be empty")
	}

	for _, order := range orders {
		amount += order.Total
	}
	return amount, nil
}

func PlaceOrder(c *customer.Customer, i []item.Item) (*order.Order, error) {
	if len(i) == 0 {
		return nil, errors.New("order items cannot be empty")
	}

	order, err := order.New(c.ID, i)
	if err != nil {
		return nil, err
	}

	err = c.IncreaseRewards(int(order.Total) / 2)
	if err != nil {
		return nil, err
	}
	return order, nil
}
