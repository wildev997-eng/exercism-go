package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	song := []string{}
	number := map[int]string{
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
	}
	start := startBottles
	for i := takeDown; i > 0; i-- {
		if i < takeDown {
			song = append(song, "")
		}
		if start-1 == 0 {
			song = append(song, fmt.Sprintf("%s green bottle hanging on the wall,", number[start]))
			song = append(song, fmt.Sprintf("%s green bottle hanging on the wall,", number[start]))
			song = append(song, "And if one green bottle should accidentally fall,")
			song = append(song, "There'll be no green bottles hanging on the wall.")
			break
		}
		song = append(song, fmt.Sprintf("%s green bottles hanging on the wall,", number[start]))
		song = append(song, fmt.Sprintf("%s green bottles hanging on the wall,", number[start]))
		song = append(song, "And if one green bottle should accidentally fall,")
		if start-1 == 1 {
			song = append(song, fmt.Sprintf("There'll be %s green bottle hanging on the wall.", strings.ToLower(number[start-1])))
		} else {
			song = append(song, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(number[start-1])))
		}

		start--
	}

	return song
}
