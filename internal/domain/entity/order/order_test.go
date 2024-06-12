package order_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/order"
	"github.com/stretchr/testify/assert"
)

func TestOrderEntity_New(t *testing.T) {
	t.Run("Should throw an error if customer ID is empty", func(t *testing.T) {
		macbook, _ := item.New("macbook m3 pro", 1999.9)
		airpods, _ := item.New("airpods 3", 299.9)

		orderItems := []item.Item{*macbook, *airpods}

		_, err := order.New("", orderItems)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "customer ID cannot be empty")
	})

	t.Run("Should throw an error if order items is empty", func(t *testing.T) {
		emptyItems := []item.Item{}

		_, err := order.New("1", emptyItems)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "order must have at least one item")
	})

	t.Run("Should calculate total correctly", func(t *testing.T) {
		macbook, _ := item.New("macbook m3 pro", 10)
		airpods, _ := item.New("airpods 3", 20)

		orderItems := []item.Item{*macbook, *airpods}

		c, err := order.New("a", orderItems)
		assert.Nil(t, err)
		assert.Equal(t, c.Total, float64(30))
	})
}
