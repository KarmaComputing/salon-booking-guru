package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// Generate a random string totalling 64 bytes in size.
func GenerateToken() (string, error) {
	c := 64
	b := make([]byte, c)

	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error: Failed to generate a token")
		log.Println(err)
		return "", err
	}

	sessionId := base64.StdEncoding.EncodeToString(b)

	return sessionId, nil
}
