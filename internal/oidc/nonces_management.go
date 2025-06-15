package oidc

import (
	"fmt"
	"time"

	"github.com/earaujoassis/space/internal/services/volatile"
	"github.com/earaujoassis/space/internal/security"
)

const (
	NonceTTL = 48 * time.Hour
	NoncePrepend = "oidc.nonce"
)

func isValidNonce(nonce string) bool {
	return security.ValidNonce(nonce)
}

func storeNonceForClient(key, nonce string) bool {
	var ok bool

	nonceKey := fmt.Sprintf("%s:%s:%s", NoncePrepend, key, nonce)
	volatile.TransactionWrapper(func() {
		ok = volatile.SetKeyWithExpiration(nonceKey, 1, NonceTTL)
	})

	return ok
}
