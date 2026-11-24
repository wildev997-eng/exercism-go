package grains

import "errors"

func Square(number int) (uint64, error) {
	if number > 64 || number <= 0 {
		return 0, errors.New("Number must be lower than 64 and bigger than 0")
	}
	grain := 1
	for i := 1; i < number; i++ {
		grain *= 2
	}
	return uint64(grain), nil
}

func Total() uint64 {
	return uint64(18446744073709551615)

}
