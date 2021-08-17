package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

func TestEncryptPassword(t *testing.T) {
	pass, err := EncryptPassword("password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)
	assert.Len(t, pass, 60)

	err = VerifyPassword(pass, "password123")
	assert.NoError(t, err)
}

func TestVerifyPassword(t *testing.T) {
	pass, err := EncryptPassword("password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, pass)

	assert.NoError(t, VerifyPassword(pass, "password123"))

	assert.Error(t, VerifyPassword(pass, pass))
	assert.Error(t, VerifyPassword(pass, "password_err"))
}
