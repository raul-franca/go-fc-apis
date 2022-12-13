package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("raul", "raul@email.com", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "raul", user.Name)
	assert.Equal(t, "raul@email.com", user.Email)

}
func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("raul", "raul@email.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
