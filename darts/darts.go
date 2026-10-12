package darts

import "math"

func Score(x, y float64) int {
	radius := 100
	where := math.Abs(x)*math.Abs(x) + math.Abs(y)*math.Abs(y)
	if where > float64(radius) {
		return 0
	} else if where <= 1 {
		return 10
	} else if where <= 25 {
		return 5
	} else {
		return 1
	}
}
