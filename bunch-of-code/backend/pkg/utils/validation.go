package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
)

var emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

func IsValidEmail(s string) bool {
	return emailRe.MatchString(s)
}

func NewID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
