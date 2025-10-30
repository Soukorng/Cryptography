package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// HashMD5 returns the lowercase hexadecimal MD5 of the input.
func HashMD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
