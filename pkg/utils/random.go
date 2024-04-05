package utils

import (
	"math/rand"
	"time"
)

func Random(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
