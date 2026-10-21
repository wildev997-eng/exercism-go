package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, chicken := range birdsPerDay {
		total += chicken
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var start, total int
	if week == 1 {
		start = 0
	} else {
		start = (week - 1) * 7
	}
	end := week * 7

	for _, chicken := range birdsPerDay[start:end] {
		total += chicken
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
			continue
		}
	}
	return birdsPerDay
}
