package repository

import (
	"fmt"
	"time"

	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/internal/models"
)

const (
	NonceTTL     = 48 * time.Hour
	NonceCodeTTL = 10 * time.Minute
	NoncePrepend = "oidc.nonce"
)

type NonceRepository struct {
	*BaseMemoryRepository[models.Nonce]
}

func NewNonceRepository(ms *memory.MemoryService) *NonceRepository {
	return &NonceRepository{
		BaseMemoryRepository: NewBaseMemoryRepository[models.Nonce](ms),
	}
}

func (r *NonceRepository) Create(nonce models.Nonce) bool {
	var ok bool

	if !nonce.IsValid() {
		return false
	}

	key := nonce.ClientKey
	code := nonce.Code
	nonceStr := nonce.Nonce

	nonceKey := fmt.Sprintf("%s:%s:%s", NoncePrepend, key, nonceStr)
	r.ms.Transaction(func(c *memory.Commands) {
		ok = c.SetKeyNXWithExpiration(nonceKey, 1, NonceTTL)
	})

	if !ok {
		return ok
	}

	if code != "" {
		codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
		r.ms.Transaction(func(c *memory.Commands) {
			ok = c.SetKeyWithExpiration(codeKey, nonceStr, NonceCodeTTL)
		})
	}

	return ok
}

func (r *NonceRepository) RetrieveByCode(code string) models.Nonce {
	var nonce string

	if code == "" {
		return models.Nonce{Code: code}
	}

	codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
	r.ms.Transaction(func(c *memory.Commands) {
		nonce = c.GetKey(codeKey).ToString()
	})

	return models.Nonce{
		Code:  code,
		Nonce: nonce,
	}
}
