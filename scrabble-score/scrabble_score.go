package scrabblescore

import (
	"strings"
)

func Score(word string) int {
	points := 0
	for _, index := range word {
		switch {
		case strings.Contains("AEIOULNRST", strings.ToUpper(string(index))):
			points += 1
		case strings.Contains("DG", strings.ToUpper(string(index))):
			points += 2
		case strings.Contains("BCMP", strings.ToUpper(string(index))):
			points += 3
		case strings.Contains("FHVWY", strings.ToUpper(string(index))):
			points += 4
		case strings.Contains("K", strings.ToUpper(string(index))):
			points += 5
		case strings.Contains("JX", strings.ToUpper(string(index))):
			points += 8
		default:
			points += 10
		}
	}
	return points
}
