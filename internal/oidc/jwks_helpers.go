package oidc

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func parsePrivateKey(block *pem.Block) (*rsa.PrivateKey, error) {
	if privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return privateKey, nil
	}

	if key, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if rsaKey, ok := key.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, fmt.Errorf("PKCS#8 key is not RSA")
	}

	return nil, fmt.Errorf("unable to parse private key")
}

func convertToBase64(b []byte) string {
	return base64.URLEncoding.
		WithPadding(base64.NoPadding).
		EncodeToString(b)
}

func initKeyManager() (*KeyManager, error) {
	km := &KeyManager{}

	err := km.LoadKeysFromPath("configs/jwks")
	if err != nil {
		return nil, err
	}

	return km, nil
}

func getJWTValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	kidInterface, exists := token.Header["kid"]
	if !exists {
		return nil, fmt.Errorf("missing kid in token header")
	}

	kid, ok := kidInterface.(string)
	if !ok {
		return nil, fmt.Errorf("invalid kid format")
	}

	publicKey, err := getPublicKeyByID(kid)
	if err != nil {
		return nil, fmt.Errorf("unknown key id: %s", kid)
	}

	return publicKey, nil
}

func getPublicKeyByID(keyID string) (*rsa.PublicKey, error) {
	keyManager, err := initKeyManager()
	if err != nil {
		logs.Propagatef(logs.Error, "JWKS is not available: %s", err)
		return nil, err
	}
	keyPair := keyManager.GetKeyByID(keyID)
	if keyPair == nil {
		return nil, fmt.Errorf("key not found: %s", keyID)
	}

	if keyPair.PublicKey == nil {
		return nil, fmt.Errorf("public key not available: %s", keyID)
	}

	return keyPair.PublicKey, nil
}

func identifyTokenType(tokenString string) shared.TokenType {
	parts := strings.Split(tokenString, ".")
	if len(parts) == 3 {
		if _, err := jwt.Parse(tokenString, getJWTValidationKey); err == nil {
			return shared.TokenTypeIDToken
		}
	}

	return shared.TokenTypeAccessToken
}

func generateJWKSETag(keys []utils.H) string {
	data, _ := json.Marshal(keys)
	hash := sha256.Sum256(data)

	return fmt.Sprintf(`"%x"`, hash[:8])
}
