package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncriptKey(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	encrypted := hash.Sum(nil)
	return hex.EncodeToString(encrypted)
}