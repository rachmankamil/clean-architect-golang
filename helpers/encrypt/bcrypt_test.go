package encrypt_test

import (
	"ca-amartha/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var password string

func TestMain(m *testing.M) {
	password = "secret"
}

func TestEncrypt(t *testing.T) {
	t.Run("Hash the password, Valid condition", func(t *testing.T) {
		_, err := encrypt.Hash(password)
		assert.Nil(t, err)
	})

	t.Run("Validate the hash, Valid condition", func(t *testing.T) {
		encr, err := encrypt.Hash(password)
		assert.Nil(t, err)

		assert.True(t, encrypt.ValidateHash(password, encr))
	})

}
