package oidc

import (
	"encoding/base64"
	"crypto/rsa"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

//lint:ignore U1000 keep it for consistency
func convertToBase64(b []byte) string {
	return base64.URLEncoding.
		WithPadding(base64.NoPadding).
		EncodeToString(b)
}

//lint:ignore U1000 keep it for consistency
func rsaToJWK(publicKey *rsa.PublicKey, keyID string) utils.H {
	return utils.H{
		"kty": "RSA",
		"use": "sig",
		"kid": keyID,
		"alg": "RS256",
		"n":   convertToBase64(publicKey.N.Bytes()),
		"e":   convertToBase64(big.NewInt(int64(publicKey.E)).Bytes()),
	}
}

func jwksHandler(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=86400")
	c.Header("ETag", "")
	c.String(http.StatusNotImplemented, "Not implemented")
}
