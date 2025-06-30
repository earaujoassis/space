package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNonceModel(t *testing.T) {
	nonce := Nonce{}

	assert.False(t, nonce.IsValid())
	assert.False(t, IsValid("validate", nonce))
	val := validateModel("validate", nonce)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))

	nonce = Nonce{
		Nonce: GenerateRandomString(7),
	}

	assert.False(t, nonce.IsValid())
	assert.False(t, IsValid("validate", nonce))
	val = validateModel("validate", nonce)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))

	nonce = Nonce{
		Nonce: GenerateRandomString(129),
	}

	assert.False(t, nonce.IsValid())
	assert.False(t, IsValid("validate", nonce))
	val = validateModel("validate", nonce)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))

	nonce = Nonce{
		Nonce: GenerateRandomString(128),
	}

	assert.True(t, nonce.IsValid())
	assert.True(t, IsValid("validate", nonce))
	val = validateModel("validate", nonce)
	assert.Nil(t, val, fmt.Sprintf("%v", val))
}
