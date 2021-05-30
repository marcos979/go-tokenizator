package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

func GenerateToken(url string) (string, error) {
	data, err := extractDataToHash(url)
	if err != nil {
		return "", err
	}

	secret := os.Getenv("TOKENIZATOR_SECRET")
	if secret == "" {
		log.Println("There is an empty secret.")
	}
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	shaToken := hex.EncodeToString(h.Sum(nil))

	return shaToken, nil
}

func GenerateTokenizedUrl(url string) (string, error) {
	token, err := GenerateToken(url)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(url)
	sb.WriteString("/token=")
	sb.WriteString(token)

	return sb.String(), nil
}

func extractDataToHash(url string) (string, error) {
	return url, nil
}
