package customer_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/valueobj/address"
	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity_New(t *testing.T) {
	t.Run("Should throw error if customer name is empty", func(t *testing.T) {
		addr := address.Address{Street: "street xpto", City: "city example", Zip: "99999"}
		_, err := customer.New("", addr)

		assert.EqualError(t, err, "customer name cannot be empty")
	})

	t.Run("Should change customer name", func(t *testing.T) {
		addr := address.Address{Street: "street xpto", City: "city example", Zip: "99999"}
		c, err := customer.New("john doe", addr)

		assert.Nil(t, err)
		assert.Equal(t, c.Name, "john doe")

		err = c.ChangeName("jane doe")

		assert.Nil(t, err)
		assert.Equal(t, c.Name, "jane doe")
		assert.NotEqual(t, c.Name, "john doe")
	})

	t.Run("Should activate a customer", func(t *testing.T) {
		addr := address.Address{Street: "street xpto", City: "city example", Zip: "99999"}
		c, _ := customer.New("john doe", addr)

		c.Enable()
		assert.True(t, c.Enabled, true)
	})

	t.Run("Should disable a customer", func(t *testing.T) {
		addr := address.Address{Street: "street xpto", City: "city example", Zip: "99999"}
		c, _ := customer.New("john doe", addr)

		c.Enable()
		assert.True(t, c.Enabled)

		c.Disable()
		assert.False(t, c.Enabled)
	})
}
