package highscores

import (
	"slices"
	"sort"
)

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return slices.Max(s.scores)
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	copies := make([]int, len(s.scores))
	copy(copies, s.scores)
	sort.Slice(copies, func(i, j int) bool {
		return copies[i] > copies[j] // Descending
	})

	if len(copies) < 3 {
		return copies[:]
	}
	return copies[:3]
}
