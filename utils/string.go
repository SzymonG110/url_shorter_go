package utils

import (
	"math/rand"
	"time"
)

func GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLengthMin = 3
	const codeLengthMax = 8
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(codeLengthMax-codeLengthMin+1) + codeLengthMin
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
