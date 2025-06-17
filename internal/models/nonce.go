package models

import (
	"github.com/earaujoassis/space/internal/security"
)

type Nonce struct {
	Key string
}

func (nonce *Nonce) IsValid() bool {
	return security.ValidNonce(nonce.Key)
}
