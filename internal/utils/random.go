package utils

import (
	"log"
	"math/rand"
)

func GenerateRandomString(size int) string {
	var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, size)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	log.Println("Created Token: ", string(b))

	return string(b)
}
