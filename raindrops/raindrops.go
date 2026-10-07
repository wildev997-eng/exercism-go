package raindrops

import "fmt"

func Convert(number int) string {
	if number%3 == 0 {
		if number%5 == 0 {
			if number%7 == 0 {
				return "PlingPlangPlong"
			}
			return "PlingPlang"
		} else if number%7 == 0 {
			return "PlingPlong"
		}
		return "Pling"
	} else if number%5 == 0 {
		if number%7 == 0 {
			return "PlangPlong"
		}
		return "Plang"
	} else if number%7 == 0 {
		return "Plong"
	} else {
		return fmt.Sprintf("%d", number)
	}
}
