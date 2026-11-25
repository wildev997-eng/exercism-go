package lineup

import (
	"fmt"
)

func Format(name string, number int) string {
	suffix := "th"

	switch {
	case number%100 >= 11 && number%100 <= 13:
		suffix = "th"
	case number%10 == 1:
		suffix = "st"
	case number%10 == 2:
		suffix = "nd"
	case number%10 == 3:
		suffix = "rd"
	}

	return fmt.Sprintf(
		"%s, you are the %d%s customer we serve today. Thank you!",
		name, number, suffix,
	)
}
