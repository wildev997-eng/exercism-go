package atbashcipher

import (
	"strings"
)

func Atbash(s string) string {
	calm := strings.ToLower(s)
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	var sb strings.Builder
	count := 0

	for _, content := range calm {
		position := strings.IndexRune(plain, content)

		if position != -1 {
			if count > 0 && count%5 == 0 {
				sb.WriteByte(' ')
			}

			sb.WriteByte(cipher[position])
			count++

		} else if content >= '0' && content <= '9' {
			if count > 0 && count%5 == 0 {
				sb.WriteByte(' ')
			}

			sb.WriteRune(content)
			count++
		}
	}
	return sb.String()
}
