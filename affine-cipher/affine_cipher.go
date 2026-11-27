package affinecipher

import (
	"errors"
	"strings"
)

func IsCoprime(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}
	return a == 1
}

func Encode(text string, a, b int) (string, error) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	var sb strings.Builder
	count := 0

	if !IsCoprime(a, 26) {
		return "", errors.New("number is not coprime")
	}

	for _, letter := range strings.ToLower(text) {
		var output byte

		switch {
		case letter >= 'a' && letter <= 'z':
			x := int(letter - 'a')
			where := ((a*x+b)%26 + 26) % 26
			output = alphabet[where]

		case letter >= '0' && letter <= '9':
			output = byte(letter)

		default:
			continue
		}

		if count > 0 && count%5 == 0 {
			sb.WriteByte(' ')
		}

		sb.WriteByte(output)
		count++
	}

	return sb.String(), nil
}

func modularInverse(a int) int {
	for i := 1; i < 26; i++ {
		if ((a*i)%26+26)%26 == 1 {
			return i
		}
	}

	return -1
}

func Decode(text string, a, b int) (string, error) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	var sb strings.Builder

	if !IsCoprime(a, 26) {
		return "", errors.New("number is not coprime")
	}

	for _, letter := range strings.ToLower(text) {
		var output byte

		switch {
		case letter >= 'a' && letter <= 'z':
			x := int(letter - 'a')
			where := ((modularInverse(a) * (x - b) % 26) + 26) % 26
			output = alphabet[where]

		case letter >= '0' && letter <= '9':
			output = byte(letter)

		default:
			continue
		}
		sb.WriteByte(output)
	}

	return sb.String(), nil
}
