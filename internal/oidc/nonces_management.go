package oidc

import (
	"fmt"
	"time"

	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services/volatile"
)

const (
	NonceTTL     = 48 * time.Hour
	NonceCodeTTL = 10 * time.Minute
	NoncePrepend = "oidc.nonce"
)

func isValidNonce(nonce string) bool {
	return security.ValidNonce(nonce)
}

func storeNonceForClient(key, nonce, code string) bool {
	var ok bool

	nonceKey := fmt.Sprintf("%s:%s:%s", NoncePrepend, key, nonce)
	volatile.TransactionWrapper(func() {
		ok = volatile.SetKeyNXWithExpiration(nonceKey, 1, NonceTTL)
	})

	if !ok {
		return ok
	}

	if code != "" {
		codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
		volatile.TransactionWrapper(func() {
			volatile.SetKeyWithExpiration(codeKey, nonce, NonceCodeTTL)
		})
	}

	return ok
}

func retrieveNonceForCode(code string) string {
	var nonce string

	codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
	volatile.TransactionWrapper(func() {
		nonce = volatile.GetKey(codeKey).ToString()
	})

	return nonce
}
