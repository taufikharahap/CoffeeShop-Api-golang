package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd12345"
var hasedPassword string
var errors error

func TestHashPassword(t *testing.T) {
	hasedPassword, errors = HashPassword(password)
	assert.NoError(t, errors, "error occured while hasing password")
	assert.NotEqual(t, password, hasedPassword, "password tidak terhasing")
}

func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var hasPassword = VerifyPassword(hasedPassword, password)
		assert.Nil(t, hasPassword, "password salah")
	})

	t.Run("verify failed", func(t *testing.T) {
		var hasPassword = VerifyPassword(hasedPassword, "12345")
		assert.NotNil(t, hasPassword, "password masih benar")
	})
}
