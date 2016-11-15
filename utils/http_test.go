package utils

import (
    "testing"
    "encoding/base64"

    "github.com/stretchr/testify/assert"
)

func TestBasicAuthEncode(t *testing.T) {
    secret := "my-secret"
    key := "my-key"
    authorization := key + ":" + secret
    encodedAuth := base64.StdEncoding.EncodeToString([]byte(authorization))
    assert.Equal(t, encodedAuth, BasicAuthEncode(key, secret), "should create correct encoded Basic Authorization")

    fakeSecret := "fake-secret"
    fakeAuthorization := key + ":" + fakeSecret
    fakeEncodedAuth := base64.StdEncoding.EncodeToString([]byte(fakeAuthorization))
    assert.NotEqual(t, fakeEncodedAuth, BasicAuthEncode(key, secret), "should create incorrect encoded Basic Authorization")
}

func TestBasicAuthDecode(t *testing.T) {
    secret := "my-secret"
    key := "my-key"
    encodedAuth := BasicAuthEncode(key, secret)
    keyDecoded, secretDecoded := BasicAuthDecode(encodedAuth)
    assert.Equal(t, key, keyDecoded, "should have returned correct key")
    assert.Equal(t, secret, secretDecoded, "should have returned correct secret")

    fakeSecret := "fake-secret"
    fakeEncodedAuth := BasicAuthEncode(key, fakeSecret)
    keyDecoded, secretDecoded = BasicAuthDecode(fakeEncodedAuth)
    assert.Equal(t, key, keyDecoded, "should have returned correct key")
    assert.NotEqual(t, secret, secretDecoded, "should not have returned correct secret")
    assert.Equal(t, fakeSecret, secretDecoded, "should have returned fake secret")
}
