package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	container := map[string]int{}
	reg := regexp.MustCompile(`[^a-zA-Z]+`)
	removedSpec := strings.ToLower(reg.ReplaceAllString(input, ""))

	for _, content := range removedSpec {
		container[string(content)]++
	}
	if len(container) < 26 {
		return false
	}
	return true
}
