package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	reccomendation := '❗'
	search := '🔍'
	weather := '☀'
	what := "default"
	for _, content := range log {
		if content == reccomendation {
			what = "recommendation"
			break
		}
		if content == search {
			what = "search"
			break
		}
		if content == weather {
			what = "weather"
			break
		}
	}
	return what
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	} else {
		return false
	}
}
