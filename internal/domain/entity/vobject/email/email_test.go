package email_test

import (
	"testing"

	"github.com/bernardinorafael/fc-ddd-pattern/internal/domain/entity/vobject/email"
	"github.com/stretchr/testify/assert"
)

func TestEmailValueObject(t *testing.T) {
	t.Run("Should throw error if email is in wrong format", func(t *testing.T) {
		_, err := email.New("john.doeemail.com")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid email format")
	})

	t.Run("Should return email format validated", func(t *testing.T) {
		address, err := email.New("john.doe@email.com")
		emailString := address.GetEmail()

		assert.Nil(t, err)
		assert.Equal(t, string(emailString), "john.doe@email.com")
	})
}
