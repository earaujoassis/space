package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidLanguageModel(t *testing.T) {
	language := Language{
		Name:    "",
		IsoCode: "",
	}

	assert.False(t, IsValid("validate", language))
	val := validateModel("validate", language)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))
	err := language.BeforeSave(nil)
	assert.NotNil(t, err, fmt.Sprintf("%v", err))

	language = Language{
		Name:    "Português (Brasil)",
		IsoCode: "pt-BR",
	}

	assert.True(t, IsValid("validate", language))
	val = validateModel("validate", language)
	assert.Nil(t, val, fmt.Sprintf("%v", val))
	err = language.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}
