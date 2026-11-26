package cipher

import "strings"

type shift int
type vigenere string

func NewCaesar() Cipher {
	return shift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}

	return shift(distance)
}

func (c shift) Encode(input string) string {
	keyLetter := byte('a' + (int(c)+26)%26)
	return transform(input, string(keyLetter), 1)
}

func (c shift) Decode(input string) string {
	return (-c).Encode(input)
}

func NewVigenere(key string) Cipher {
	hasNonA := false

	for _, letter := range key {
		if letter < 'a' || letter > 'z' {
			return nil
		}

		if letter != 'a' {
			hasNonA = true
		}
	}

	if !hasNonA {
		return nil
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	return transform(input, string(v), 1)
}

func (v vigenere) Decode(input string) string {
	return transform(input, string(v), -1)
}

func transform(input, key string, direction int) string {
	var output strings.Builder
	output.Grow(len(input))

	for i := 0; i < len(input); i++ {
		letter := input[i]
		if letter >= 'A' && letter <= 'Z' {
			letter += 'a' - 'A'
		}
		if letter < 'a' || letter > 'z' {
			continue
		}

		distance := int(key[output.Len()%len(key)] - 'a')
		position := int(letter-'a') + direction*distance
		output.WriteByte(byte('a' + (position+26)%26))
	}

	return output.String()
}
