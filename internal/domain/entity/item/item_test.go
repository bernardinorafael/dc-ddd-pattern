package item_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/item"
	"github.com/stretchr/testify/assert"
)

func TestOrderItem_New(t *testing.T) {
	t.Run("Should create a new order item", func(t *testing.T) {
		orderItem, err := item.New("testing order item", 100.0)

		assert.Nil(t, err)
		assert.Equal(t, orderItem.Name, "testing order item")
		assert.Equal(t, orderItem.Price, 100.0)
	})
}
