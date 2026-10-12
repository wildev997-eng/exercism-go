package collatzconjecture

import "errors"

func CollatzConjecture(num int) (int, error) {
	i := 0
	if num <= 0 {
		return 0, errors.New("Numbers must be positive.")
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		i++

	}
	return i, nil
}
