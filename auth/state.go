package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func StateGenerator() (string, error) {
	// Generate a random byte slice
	randBytes := make([]byte, 32)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "[Auth: Generating state]", err
	}

	// Encode the random bytes to base64
	state := base64.URLEncoding.EncodeToString(randBytes)

	return state, nil
}
