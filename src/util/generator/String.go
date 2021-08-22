package generator

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	createdString := make([]byte, n)
	for index := range createdString {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		createdString[index] = letters[num.Int64()]
	}

	return string(createdString), nil
}
