package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// HashSHA1 returns the lowercase hexadecimal SHA-1 of the input.
func HashSHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
