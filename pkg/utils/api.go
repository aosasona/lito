package utils

import (
	"crypto/rand"
	"io"
)

var allowedChars = []rune("abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ123456789!=_-")

func GenerateAPIKey() string {
	return GenerateRandomString(32)
}

func GenerateRandomString(length int) string {
	str := make([]byte, length)

	n, err := io.ReadAtLeast(rand.Reader, str, length)
	if n != length || err != nil {
		return ""
	}

	for i := 0; i < length; i++ {
		str[i] = byte(allowedChars[int(str[i])%len(allowedChars)])
	}

	return string(str)
}
