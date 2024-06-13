package customer_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/customer"
	"github.com/stretchr/testify/assert"
)

func TestCustomerEntity_New(t *testing.T) {
	t.Run("Should throw error if customer name is empty", func(t *testing.T) {
		_, err := customer.New("", "john.doe@email.com")

		assert.EqualError(t, err, "customer name cannot be empty")
	})

	t.Run("Should change customer name", func(t *testing.T) {
		c, err := customer.New("john doe", "john.doe@email.com")

		assert.Nil(t, err)
		assert.Equal(t, c.Name, "john doe")

		err = c.ChangeName("jane doe")

		assert.Nil(t, err)
		assert.Equal(t, c.Name, "jane doe")
		assert.NotEqual(t, c.Name, "john doe")
	})

	t.Run("Should activate a customer", func(t *testing.T) {
		c, _ := customer.New("john doe", "john.doe@email.com")

		c.Enable()
		assert.True(t, c.Enabled, true)
	})

	t.Run("Should disable a customer", func(t *testing.T) {
		c, _ := customer.New("john doe", "john.doe@email.com")

		c.Enable()
		assert.True(t, c.Enabled)

		c.Disable()
		assert.False(t, c.Enabled)
	})

	t.Run("Should increase customer rewards correctly", func(t *testing.T) {
		c, err := customer.New("john doe", "john.doe@email.com")

		assert.Nil(t, err)
		assert.Equal(t, c.Rewards, 0)

		err = c.IncreaseRewards(10)
		assert.Nil(t, err)
		assert.Equal(t, c.Rewards, 10)
	})

	t.Run("Should throw an error if rewards points is lesser than zero", func(t *testing.T) {
		c, _ := customer.New("john doe", "john.doe@email.com")

		err := c.IncreaseRewards(0)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "rewards points must be greater than zero")
	})
}
