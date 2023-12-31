package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func DigestSHA256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
