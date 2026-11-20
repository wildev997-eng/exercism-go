// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var abb strings.Builder
	cleaned := strings.Map(func(r rune) rune {
		if r == '-' {
			return ' '
		}

		if unicode.IsPunct(r) {
			return -1
		}

		return r
	}, s)

	words := strings.Fields(cleaned)
	for _, content := range words {
		abb.WriteString(string(content[0]))
	}
	return strings.ToUpper(abb.String())
}
