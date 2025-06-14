package oidc

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type KeyPair struct {
	ID         string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	CreatedAt  time.Time
}

type KeyManager struct {
	Keys []KeyPair
}

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

func (km *KeyManager) loadPrivateKey(path string) error {
	keyID := km.extractKeyID(path, ".private.pem")

	keyData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return fmt.Errorf("failed to parse PEM block from %s", path)
	}

	privateKey, err := parsePrivateKey(block)
	if err != nil {
		return err
	}

	keyPair := km.findOrCreateKeyPair(keyID)
	keyPair.PrivateKey = privateKey
	keyPair.PublicKey = &privateKey.PublicKey

	return nil
}

func (km *KeyManager) loadPublicKey(path string) error {
	keyID := km.extractKeyID(path, ".public.pem")

	keyData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return fmt.Errorf("failed to parse PEM block from %s", path)
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return fmt.Errorf("not an RSA public key: %s", path)
	}

	keyPair := km.findOrCreateKeyPair(keyID)
	keyPair.PublicKey = rsaPublicKey

	return nil
}

func (km *KeyManager) extractKeyID(path, suffix string) string {
	filename := filepath.Base(path)
	return strings.TrimSuffix(filename, suffix)
}

func (km *KeyManager) findOrCreateKeyPair(keyID string) *KeyPair {
	for i := range km.Keys {
		if km.Keys[i].ID == keyID {
			return &km.Keys[i]
		}
	}

	newKeyPair := KeyPair{
		ID:        keyID,
		CreatedAt: time.Now(),
	}
	km.Keys = append(km.Keys, newKeyPair)

	return &km.Keys[len(km.Keys)-1]
}

func convertToBase64(b []byte) string {
	return base64.URLEncoding.
		WithPadding(base64.NoPadding).
		EncodeToString(b)
}
