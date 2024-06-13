package domainservice_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/domainservice"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/order"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/product"
	"github.com/stretchr/testify/assert"
)

func TestDomainService_IncProductPercent(t *testing.T) {
	t.Run("Should update products price", func(t *testing.T) {
		iphone, _ := product.New("iphone 13 pro max", 100.0, 1)
		macbook, _ := product.New("macbook pro m3 pro", 100.0, 1)
		airpods, _ := product.New("airpods 3", 100.0, 1)

		products := []*product.Product{iphone, macbook, airpods}

		err := domainservice.IncProductPercent(products, 100)

		assert.Nil(t, err)
		assert.Equal(t, iphone.Price, 200.0)
		assert.Equal(t, macbook.Price, 200.0)
		assert.Equal(t, airpods.Price, 200.0)
	})

	t.Run("Should throw error if products slice is empty", func(t *testing.T) {
		err := domainservice.IncProductPercent([]*product.Product{}, 15)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "products slice cannot be empty")
	})

	t.Run("Should throw error if percent value is greater than 100", func(t *testing.T) {
		iphone, _ := product.New("iphone 13 pro max", 100.0, 1)
		macbook, _ := product.New("macbook pro m3 pro", 100.0, 1)
		airpods, _ := product.New("airpods 3", 100.0, 1)

		products := []*product.Product{iphone, macbook, airpods}

		err := domainservice.IncProductPercent(products, 101)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "percent value must be between 1 and 100")
	})
}

func TestDomainService_GetOrdersAmount(t *testing.T) {
	t.Run("Should return the orders amount correctly", func(t *testing.T) {
		iphone, _ := item.New("iphone 13 pro max", 100.0)
		macbook, _ := item.New("macbook pro m3 pro", 100.0)
		airpods, _ := item.New("airpods 3", 100.0)

		order1, _ := order.New("1", []item.Item{*iphone, *macbook, *airpods})
		order2, _ := order.New("1", []item.Item{*iphone, *macbook})

		amount, err := domainservice.GetOrdersAmount([]order.Order{*order1, *order2})

		assert.Nil(t, err)
		assert.Equal(t, amount, 500.0)
	})

	t.Run("Should throw an error if orders slice is empty", func(t *testing.T) {
		_, err := domainservice.GetOrdersAmount([]order.Order{})

		assert.NotNil(t, err)
		assert.EqualError(t, err, "orders slice cannot be empty")
	})
}

func TestDomainService_PlaceOrder(t *testing.T) {
	t.Run("Should throw an error if order items is empty", func(t *testing.T) {
		c, _ := customer.New("john doe", "john.doe@email.com")
		_, err := domainservice.PlaceOrder(c, []item.Item{})

		assert.NotNil(t, err)
		assert.EqualError(t, err, "order items cannot be empty")
	})

	t.Run("Should place a new order correctly", func(t *testing.T) {
		iphone, _ := item.New("iphone 13 pro max", 100.0)
		macbook, _ := item.New("macbook pro m3 pro", 100.0)

		c, _ := customer.New("john doe", "john.doe@email.com")
		order, err := domainservice.PlaceOrder(c, []item.Item{*iphone, *macbook})

		assert.Nil(t, err)
		assert.Equal(t, order.Total, 200.0)
		assert.Equal(t, c.Rewards, 100)
	})
}
