package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	anagram := []string{}
	for _, sub := range candidates {
		if strings.ToLower(sub) == strings.ToLower(subject) {
			continue
		} else if sortWord(subject) == sortWord(sub) {
			anagram = append(anagram, sub)
		}
	}
	return anagram
}

func sortWord(word string) string {
	letters := []rune(strings.ToLower(word))
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}
