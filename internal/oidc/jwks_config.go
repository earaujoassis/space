package oidc

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/earaujoassis/space/internal/utils"
)

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

func GenerateJWKSETag(keys []utils.H) string {
	data, _ := json.Marshal(keys)
	hash := sha256.Sum256(data)

	return fmt.Sprintf(`"%x"`, hash[:8])
}
