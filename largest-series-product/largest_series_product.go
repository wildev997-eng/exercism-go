package largestseriesproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span cannot be negative")
	}

	if span > len(digits) {
		return 0, errors.New("span exceeds string length")
	}

	for _, digit := range digits {
		if digit < '0' || digit > '9' {
			return 0, errors.New("invalid digit")
		}
	}

	if span == 0 {
		return 1, nil
	}

	var biggest int64

	for i := 0; i <= len(digits)-span; i++ {
		total := int64(1)

		for _, digit := range digits[i : i+span] {
			total *= int64(digit - '0')
		}

		if total > biggest {
			biggest = total
		}
	}

	return biggest, nil
}
