package shatool

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256String finish following steps:
// 1. input string => []byte
// 2. calcu sha256
// 3. convert result from []byte to string
func Sha256String(input string) (output string) {
	return hex.EncodeToString(
		sha256.New().Sum([]byte(input)),
	)
}
