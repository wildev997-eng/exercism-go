package dndcharacter

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	result := math.Floor(float64(score-10) / 2)
	return int(result)
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return randomNumber()
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}

func randomNumber() int {
	random := []int{}
	for i := 4; i > 0; i-- {
		num := rand.IntN(6) + 1
		random = append(random, num)
	}
	slices.SortFunc(random, func(a, b int) int {
		return b - a
	})
	random = random[:len(random)-1]
	num := 0
	for _, stats := range random {
		num += stats
	}
	return num

}
