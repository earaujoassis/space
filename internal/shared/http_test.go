package shared

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthEncode(t *testing.T) {
	secret := "my-secret"
	key := "my-key"
	authorization := key + ":" + secret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authorization))
	assert.Equal(t, encodedAuth, BasicAuthEncode(key, secret))

	fakeSecret := "fake-secret"
	fakeAuthorization := key + ":" + fakeSecret
	fakeEncodedAuth := base64.StdEncoding.EncodeToString([]byte(fakeAuthorization))
	assert.NotEqual(t, fakeEncodedAuth, BasicAuthEncode(key, secret))
}

func TestBasicAuthDecode(t *testing.T) {
	secret := "my-secret"
	key := "my-key"
	encodedAuth := BasicAuthEncode(key, secret)
	keyDecoded, secretDecoded := BasicAuthDecode(encodedAuth)
	assert.Equal(t, key, keyDecoded)
	assert.Equal(t, secret, secretDecoded)

	fakeSecret := "fake-secret"
	fakeEncodedAuth := BasicAuthEncode(key, fakeSecret)
	keyDecoded, secretDecoded = BasicAuthDecode(fakeEncodedAuth)
	assert.Equal(t, key, keyDecoded)
	assert.NotEqual(t, secret, secretDecoded)
	assert.Equal(t, fakeSecret, secretDecoded)
}

func TestBasicAuthDecodeWithNonEncodedData(t *testing.T) {
	keyDecoded, secretDecoded := BasicAuthDecode("wrongstring")
	assert.Equal(t, keyDecoded, "")
	assert.Equal(t, secretDecoded, "")
}

func TestBasicAuthDecodeWithEncodedDataWithoutAColon(t *testing.T) {
	encodedString := base64.StdEncoding.EncodeToString([]byte("wrongstring"))
	keyDecoded, secretDecoded := BasicAuthDecode(encodedString)
	assert.Equal(t, keyDecoded, "")
	assert.Equal(t, secretDecoded, "")
}

func TestMustServeJSON(t *testing.T) {
	assert.True(t, MustServeJSON("/api/users/instropect", ""))
	assert.True(t, MustServeJSON("/api/testing-only", ""))
	assert.False(t, MustServeJSON("/oauth/authorize", ""))
	assert.True(t, MustServeJSON("/oauth/token", ""))
	assert.False(t, MustServeJSON("/oidc/authorize", ""))
	assert.True(t, MustServeJSON("/oidc/token", ""))
}
