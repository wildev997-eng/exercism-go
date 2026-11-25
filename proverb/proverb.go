package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	if rhyme[0] == "nail" && len(rhyme) == 1 {
		return []string{"And all for the want of a nail."}
	}

	song := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		song = append(song, line)
	}
	song = append(song, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return song
}
