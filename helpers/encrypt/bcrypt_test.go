package encrypt_test

import (
	"ca-amartha/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "secret"
	hash, err := encrypt.Hash(password)
	assert.Nil(t, err)

	assert.True(t, encrypt.ValidateHash(password, hash))
}
