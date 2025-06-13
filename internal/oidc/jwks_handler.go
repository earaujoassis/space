package oidc

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/logs"
)

func initKeyManager() (*KeyManager, error) {
    km := &KeyManager{}

    err := km.LoadKeysFromPath("configs/jwks")
    if err != nil {
        return nil, err
    }

    return km, nil
}

func jwksHandler(c *gin.Context) {
	keyManager, err := initKeyManager()
	if err != nil {
		logs.Propagatef(logs.Error, "JWKS is not available: %s", err)
		c.JSON(http.StatusInternalServerError, utils.H{
			"error":             shared.ServerError,
			"error_description": "JWKS is not available",
		})
		return
	}
	publicKeys := keyManager.GetPublicKeys()
	c.Header("Cache-Control", "public, max-age=86400")
	c.Header("ETag", GenerateJWKSETag(publicKeys))
	c.JSON(http.StatusOK, publicKeys)
}
