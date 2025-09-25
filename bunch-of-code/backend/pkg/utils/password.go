package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(pw string) string {
	h := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(h[:])
}

func CheckPassword(hash, pw string) bool {
	return HashPassword(pw) == hash
}
