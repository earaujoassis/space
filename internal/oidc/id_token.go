package oidc

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//lint:ignore U1000 keep it for consistency
func createIDToken(issuer, userPublicId, clientKey string, privateKey *rsa.PrivateKey) string {
    claims := jwt.MapClaims{
        "iss": issuer,
        "sub": userPublicId,
        "aud": clientKey,
        "exp": time.Now().Add(time.Hour).Unix(),
        "iat": time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    token.Header["kid"] = "key-name"

    signedToken, err := token.SignedString(privateKey)
	if err != nil {
        return ""
    }
    return signedToken
}
