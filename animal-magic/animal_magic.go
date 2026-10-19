package chance

import "math/rand/v2"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.IntN(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animal := []string{
		"ant", "dog", "beaver", "cat", "elephant", "fox", "giraffe", "hedgehog",
	}
	rand.Shuffle(len(animal), func(i, j int) {
		animal[i], animal[j] = animal[j], animal[i]
	})
	return animal
}
