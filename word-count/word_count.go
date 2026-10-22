package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	howMany := map[string]int{}
	lower := strings.Clone(strings.ToLower(phrase))
	words := strings.FieldsFunc(lower, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	})

	for _, content := range words {
		content = strings.Trim(content, "'")

		if content != "" {
			howMany[content]++
		}
	}

	return howMany
}
