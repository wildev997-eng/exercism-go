package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var where int
	var x int
	id = strings.ReplaceAll(id, " ", "")

	if len(id)%2 == 0 {
		where = 0
	} else {
		where = 1
	}
	for i := 0; i < len(id); i++ {
		if i == where {
			num, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return false
			} else {
				if num*2 > 9 {
					num = num*2 - 9
					x += num
					where += 2
				} else {
					x += num * 2
					where += 2
				}
			}
		} else {
			num, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return false
			} else {
				x += num
			}
		}
	}

	if x%10 == 0 && len(id) > 1 {
		return true
	} else {
		return false
	}

}
