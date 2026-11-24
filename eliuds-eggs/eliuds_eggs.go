package eliudseggs

import "strconv"

func EggCount(displayValue int) int {
	binary := strconv.FormatInt(int64(displayValue), 2)
	howMany := 0
	for _, content := range binary {
		if content == '1' {
			howMany++
		}
	}
	return howMany
}
