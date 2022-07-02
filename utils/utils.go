package utils

import (
	"crypto/sha256"
	"fmt"
)

// Sha256sum will hash the file and return the sha256sum
func Sha256sum(data string) (string, error) {
	dataBytes := []byte(data)
	h := sha256.New()

	return fmt.Sprintf("%x", h.Sum(dataBytes)), nil
}
