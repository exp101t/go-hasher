package hash

import (
	"crypto/sha256"
)

func hashWithSha256(data []byte) []byte {
	var rawHash = sha256.Sum256(data)
	return rawHash[:]
}
