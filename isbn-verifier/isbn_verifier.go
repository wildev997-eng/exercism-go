package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	var x, sum int
	if len(strings.ReplaceAll(isbn, "-", "")) != 10 || strings.Contains(isbn, "X") && strings.Index(isbn, "X") != len(isbn)-1 {
		return false
	}
	for _, char := range isbn {
		if char == '-' {
			continue
		} else if char == 'X' {
			x++
			sum += 10 * x
		} else {
			x++
			sum += int(char-'0') * x
		}
	}
	return sum%11 == 0
}
