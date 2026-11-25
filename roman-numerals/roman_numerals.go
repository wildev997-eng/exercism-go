package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("Number must be greater than 0 and lower than 3999")
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder

	for i := 0; i < len(values); i++ {
		for input >= values[i] {
			result.WriteString(symbols[i])
			input -= values[i]
		}
	}

	return result.String(), nil

}
