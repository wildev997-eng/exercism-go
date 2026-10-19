package bafflingbirthdays

import (
	"math/rand/v2"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	same := false
	for i := 0; i < len(dates); i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].Month() == dates[j].Month() && dates[i].Day() == dates[j].Day() {
				same = true
			}
		}
	}
	return same
}

func RandomBirthdates(size int) []time.Time {
	bday := []time.Time{}
	for i := 0; i < size; i++ {
		day := rand.IntN(365)
		birthday := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, day)
		bday = append(bday, birthday)
	}
	return bday
}

func EstimatedProbability(size int) float64 {
	const trials = 10000
	matches := 0

	for i := 0; i < trials; i++ {
		dates := RandomBirthdates(size)

		if SharedBirthday(dates) {
			matches++
		}
	}

	return (float64(matches) / float64(trials)) * 100
}
