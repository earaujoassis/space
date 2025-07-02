package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	message := "Two wrongs don't make a right"
	fakeKey := []byte("Y2NvtbNymWRUUnYQ")
	key := []byte("m36mh39DtwvndHtY")
	encrypted, _ := Encrypt(key, []byte(message))
	decrypted, _ := Decrypt(key, encrypted)
	fakeDecrypted, _ := Decrypt(fakeKey, encrypted)
	assert.NotEqual(t, message, string(encrypted))
	assert.Equal(t, message, string(decrypted))
	assert.NotEqual(t, message, string(fakeDecrypted))
}
