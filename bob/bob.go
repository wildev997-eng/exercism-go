// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)

	isQuestion := strings.HasSuffix(trimmed, "?")

	isYelling := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark

	switch {
	case trimmed == "":
		return "Fine. Be that way!"

	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"

	case isQuestion:
		return "Sure."

	case isYelling:
		return "Whoa, chill out!"

	default:
		return "Whatever."
	}

}
