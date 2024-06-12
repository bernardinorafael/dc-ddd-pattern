package product_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/product"
	"github.com/stretchr/testify/assert"
)

func TestProductEntity_New(t *testing.T) {
	t.Run("Should create a product correctly", func(t *testing.T) {
		p, err := product.New("product-1", 10.9, 1)

		assert.Nil(t, err)
		assert.Equal(t, p.Name, "product-1")
		assert.Equal(t, p.Price, 10.9)
	})

	t.Run("Should thrown an error if quantity is less than zero", func(t *testing.T) {
		_, err := product.New("product-1", 10.9, 0)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "quantity must be greather than zero")
	})

	t.Run("Should change price successfully", func(t *testing.T) {
		p, err := product.New("product-1", 10.9, 1)
		assert.Nil(t, err)

		err = p.ChangePrice(100)
		assert.Nil(t, err)

		assert.Equal(t, p.Price, 100.0)
	})
}
