package models

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidSettingModel(t *testing.T) {
	setting := Setting{}
	assert.False(t, IsValid("validate", setting))
	val := validateModel("validate", setting)
	assert.NotNil(t, val, fmt.Sprintf("%v", val))
	err := setting.BeforeSave(nil)
	assert.NotNil(t, err, fmt.Sprintf("%v", err))

	client := Client{
		Name:         "internal",
		Secret:       GenerateRandomString(64),
		CanonicalURI: []string{"localhost"},
		RedirectURI:  []string{"/"},
		Scopes:       PublicScope,
		Type:         PublicClient,
	}
	err = client.BeforeSave(nil)
	require.Nil(t, err, fmt.Sprintf("%s", err))
	user := User{
		FirstName:     gofakeit.FirstName(),
		LastName:      gofakeit.LastName(),
		Username:      gofakeit.Username(),
		Email:         gofakeit.Email(),
		Passphrase:    gofakeit.Password(true, true, true, true, false, 10),
		CodeSecret:    gofakeit.Password(true, true, true, true, false, 64),
		RecoverSecret: gofakeit.Password(true, true, true, true, false, 64),
	}
	user.Client = client
	user.Language = Language{
		Name:    "English",
		IsoCode: "en-US",
	}
	err = user.BeforeSave(nil)
	require.Nil(t, err, fmt.Sprintf("%s", err))
	setting = Setting{
		User:     user,
		Realm:    "testing",
		Category: "testing-category",
		Property: "field",
		Type:     "bool",
		Value:    "false",
	}

	assert.True(t, IsValid("validate", setting))
	val = validateModel("validate", setting)
	assert.Nil(t, val, fmt.Sprintf("%v", val))
	err = setting.BeforeSave(nil)
	assert.Nil(t, err, fmt.Sprintf("%v", err))
}

func TestSettingModelMarshalJSONBool(t *testing.T) {
	setting := Setting{
		Realm:    "testing",
		Category: "testing-category",
		Property: "field",
		Type:     "bool",
		Value:    "false",
	}

	key, value := setting.Reduce()
	assert.Equal(t, "testing.testing-category.field", key)
	assert.Equal(t, false, value)
}

func TestSettingModelMarshalJSONInt(t *testing.T) {
	setting := Setting{
		Realm:    "testing",
		Category: "testing-category",
		Property: "field",
		Type:     "int",
		Value:    "10",
	}

	key, value := setting.Reduce()
	assert.Equal(t, "testing.testing-category.field", key)
	assert.Equal(t, int64(10), value)
}

func TestSettingModelMarshalJSONString(t *testing.T) {
	setting := Setting{
		Realm:    "testing",
		Category: "testing-category",
		Property: "field",
		Type:     "string",
		Value:    "10",
	}

	key, value := setting.Reduce()
	assert.Equal(t, "testing.testing-category.field", key)
	assert.Equal(t, "10", value)
}

func TestSettingModelReduceDefault(t *testing.T) {
	setting := Setting{
		Realm:    "testing",
		Category: "testing-category",
		Property: "field",
		Type:     "default",
		Value:    "default",
	}

	key, value := setting.Reduce()
	assert.Equal(t, "testing.testing-category.field", key)
	assert.Equal(t, "default", value)
}
