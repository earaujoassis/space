package repository

import (
	"fmt"
	"time"

	"github.com/earaujoassis/space/internal/gateways/redis"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
)

const (
	NonceTTL     = 48 * time.Hour
	NonceCodeTTL = 10 * time.Minute
	NoncePrepend = "oidc.nonce"
)

type NonceRepository struct {
	*BaseMemoryRepository[models.Nonce]
}

func NewNonceRepository(ms *redis.MemoryService) *NonceRepository {
	return &NonceRepository{
		BaseMemoryRepository: NewBaseMemoryRepository[models.Nonce](ms),
	}
}

func (r *NonceRepository) IsValid(nonce string) bool {
	return security.ValidNonce(nonce)
}

func (r *NonceRepository) StoreForClient(key, nonce, code string) bool {
	var ok bool

	nonceKey := fmt.Sprintf("%s:%s:%s", NoncePrepend, key, nonce)
	r.ms.Transaction(func(c *redis.Commands) {
		ok = c.SetKeyNXWithExpiration(nonceKey, 1, NonceTTL)
	})

	if !ok {
		return ok
	}

	if code != "" {
		codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
		r.ms.Transaction(func(c *redis.Commands) {
			c.SetKeyWithExpiration(codeKey, nonce, NonceCodeTTL)
		})
	}

	return ok
}

func (r *NonceRepository) RetrieveByCode(code string) string {
	var nonce string

	codeKey := fmt.Sprintf("oidc.code.nonce:%s", code)
	r.ms.Transaction(func(c *redis.Commands) {
		nonce = c.GetKey(codeKey).ToString()
	})

	return nonce
}
