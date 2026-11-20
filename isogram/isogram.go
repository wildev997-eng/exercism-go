package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	container := map[string]int{}
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	removedSpec := strings.ToLower(reg.ReplaceAllString(word, ""))

	for _, content := range removedSpec {
		container[string(content)]++
		if container[string(content)] == 2 {
			return false
		}
	}
	return true
}
