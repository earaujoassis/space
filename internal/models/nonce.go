package models

import (
	"github.com/earaujoassis/space/internal/security"
)

type Nonce struct {
	ClientKey string
	Code      string
	Nonce     string `validate:"required,min=8,max=128"`
}

func (nonce *Nonce) IsValid() bool {
	return security.ValidNonce(nonce.Nonce)
}
