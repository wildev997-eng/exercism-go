package rotationalcipher

import (
	"strings"
)

func RotationalCipher(plain string, shiftKey int) string {
	//lowercase
	plainAlpha := "abcdefghijklmnopqrstuvwxyz"
	plainAlphaUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	grabber := plainAlpha[0:shiftKey]
	grabberUpper := plainAlphaUpper[0:shiftKey]
	cipher := plainAlpha[shiftKey:]
	cipherUpper := plainAlphaUpper[shiftKey:]
	cipher += grabber
	cipherUpper += grabberUpper

	var sb strings.Builder

	for _, content := range plain {
		position := strings.IndexRune(plainAlpha, content)
		positionUp := strings.IndexRune(plainAlphaUpper, content)

		if position != -1 {
			sb.WriteByte(cipher[position])
		} else if positionUp != -1 {
			sb.WriteByte(cipherUpper[positionUp])
		} else {
			sb.WriteRune(content)
		}
	}
	return sb.String()

}
