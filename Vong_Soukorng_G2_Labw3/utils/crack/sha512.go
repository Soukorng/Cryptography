package crack

import (
	"crypto/sha512"
	"fmt"
)

// HashSHA512 returns the lowercase hex-encoded SHA-512 digest of the input string.
func HashSHA512(s string) string {
	sum := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", sum)
}
