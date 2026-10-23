package swiftscheduling

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DeliveryDate(start, delivery string) string {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, start)

	if err != nil {
		fmt.Println(err)
	}

	switch {
	case delivery == "NOW":
		t = t.Add(2 * time.Hour)
	case delivery == "ASAP":
		if t.Hour() < 13 {
			t = time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, t.Location())
		} else {
			t = time.Date(t.Year(), t.Month(), t.Day()+1, 13, 0, 0, 0, t.Location())
		}
	case delivery == "EOW":
		if t.Weekday() <= 3 {
			for t.Weekday() != 5 {
				t = t.AddDate(0, 0, 1)
			}
			t = time.Date(t.Year(), t.Month(), t.Day(), 17, 0, 0, 0, t.Location())
		} else {
			for t.Weekday() != 0 {
				t = t.AddDate(0, 0, 1)
			}
			t = time.Date(t.Year(), t.Month(), t.Day(), 20, 0, 0, 0, t.Location())
		}

	case strings.Contains(delivery, "M") == true:
		monthNumber, err := strconv.Atoi(strings.TrimSuffix(delivery, "M"))
		if err != nil {
			fmt.Println(err)
		}
		targetMonth := time.Month(monthNumber)
		targetYear := t.Year()

		if t.Month() >= targetMonth {
			targetYear++
		}

		targetStart := time.Date(targetYear, targetMonth, 1, 8, 0, 0, 0, t.Location())
		if targetStart.Weekday() > 0 && targetStart.Weekday() <= 5 {
			t = targetStart
		} else {
			for targetStart.Day() < 8 {
				if targetStart.Weekday() > 0 && targetStart.Weekday() <= 5 {
					t = targetStart
					break
				}
				targetStart = targetStart.AddDate(0, 0, 1)
			}
		}

	case strings.Contains(delivery, "Q") == true:
		qNumber, err := strconv.Atoi(strings.TrimPrefix(delivery, "Q"))
		if err != nil {
			fmt.Println(err)
		}
		endQuarter := time.Month(qNumber * 3)
		targetYear := t.Year()

		if t.Month() > endQuarter {
			targetYear++
		}

		targetStart := time.Date(targetYear, endQuarter+1, 0, 8, 0, 0, 0, time.UTC)
		for {
			if targetStart.Weekday() > 0 && targetStart.Weekday() <= 5 {
				t = targetStart
				break
			}
			targetStart = targetStart.AddDate(0, 0, -1)
		}

	}
	return t.Format("2006-01-02T15:04:05")
}
