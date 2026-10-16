package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	sorted := map[string]int{}

	for score, index := range in {
		for _, letters := range index {
			sorted[strings.ToLower(letters)] = score
		}
	}
	return sorted
}
