package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	howManyWeek, when := 0, 0
	weekCounter := 0

	switch {
	case wSched == "first":
		howManyWeek = 1
	case wSched == "second":
		howManyWeek = 2
	case wSched == "third":
		howManyWeek = 3
	case wSched == "fourth":
		howManyWeek = 4
	}

	if howManyWeek > 0 {
		for t.Day() <= lastDayOfMonth.Day() {
			if t.Weekday() == wDay {
				weekCounter++
				if weekCounter == howManyWeek {
					when = t.Day()
					break
				}
			}
			t = t.AddDate(0, 0, 1)
		}

	} else if wSched == "teenth" {
		for t.Day() <= 19 {
			if t.Day() >= 13 && t.Weekday() == wDay {
				when = t.Day()
				break
			}
			t = t.AddDate(0, 0, 1)
		}
	} else if wSched == "last" {
		t = t.AddDate(0, 0, (lastDayOfMonth.Day() - 7))
		for t.Day() <= lastDayOfMonth.Day() {
			if t.Weekday() == wDay {
				when = t.Day()
				break
			}
			t = t.AddDate(0, 0, 1)
		}
	}

	return when
}
