package oidc

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/fs"
	"os"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"github.com/earaujoassis/space/internal/utils"
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

func (km *KeyManager) LoadKeysFromPath(keysPath string) error {
	return filepath.WalkDir(keysPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".pem") {
			return nil
		}

		if strings.Contains(path, ".private.") {
			return km.loadPrivateKey(path)
		} else if strings.Contains(path, ".public.") {
			return km.loadPublicKey(path)
		}

		return nil
	})
}

func (km *KeyManager) GetPublicKeys() []utils.H {
	var publicKeys []utils.H

	for _, keyPair := range km.Keys {
		if keyPair.PublicKey == nil {
			continue
		}

		jwk := utils.H{
			"alg": "RS256",
			"e":   convertToBase64(big.NewInt(int64(keyPair.PublicKey.E)).Bytes()),
			"kid": keyPair.ID,
			"kty": "RSA",
			"n":   convertToBase64(keyPair.PublicKey.N.Bytes()),
			"use": "sig",
		}

		publicKeys = append(publicKeys, jwk)
	}

	return publicKeys
}

func (km *KeyManager) GetKeyByID(keyID string) *KeyPair {
	for i := range km.Keys {
		if km.Keys[i].ID == keyID {
			return &km.Keys[i]
		}
	}

	return nil
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
