package cipher

import (
	"strings"
)

func EncryptCaesar(text string) string {
	message := strings.ToUpper(text)
	cipherText := ""

	for _, c := range message {
		if c >= 'A' && c <= 'Z' {
			c = (c + 3) % 26
			cipherText += string(c)
		} else {
			cipherText += string(c)
		}
	}

	return cipherText
}

func DecryptCaesar(text string) string {
	cipherText := strings.ToUpper(text)
	message := ""

	for _, c := range cipherText {
		if c >= 'A' && c <= 'Z' {
			c = (c - 3) % 26
			message += string(c)
		} else {
			message += string(c)
		}
	}

	return message
}
