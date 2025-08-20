// Package cipher is for decrypting/encrypting text messages.
package cipher

import (
	"strings"
)

// CipherVigener returns encrypted text
func CipherVigenere(text, key string) string {
	if key == "" {
		return text
	}
	message := strings.ReplaceAll(strings.ToUpper(text), " ", "")
	keyword := strings.ToUpper(key)
	cipherText := ""
	keyIndex := 0

	for _, c := range message {
		if c >= 'A' && c <= 'Z' {
			c = c - 'A'
			k := rune(keyword[keyIndex]) - 'A'

			p := (c+k)%26 + 'A'
			cipherText += string(p)

			keyIndex++
			keyIndex = keyIndex % len(keyword)
		} else {
			cipherText += string(c)
		}

	}

	return cipherText
}

// DecipherVigenere returns decrypted text
func DecipherVigenere(text, key string) string {
	if key == "" {
		return text
	}
	cipherText := text
	keyword := strings.ToUpper(key)
	message := ""
	keyIndex := 0

	for _, c := range cipherText {
		if c >= 'A' && c <= 'Z' {
			c = c - 'A'
			k := rune(keyword[keyIndex]) - 'A'

			p := (c-k+26)%26 + 'A'
			message += string(p)

			keyIndex++
			keyIndex = keyIndex % len(keyword)
		} else {
			message += string(c)
		}

	}

	return message
}
