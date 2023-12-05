package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func TokenGenerate() (string, error) {
	tokenLength := 16
	tokenBytes := make([]byte, tokenLength)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(tokenBytes)

	return token, nil
}
