package squareroot

func SquareRoot(number int) (int, error) {
	result := 0
	for i := 1; i <= number; i++ {
		if i*i == number {
			result = i
			break
		}
	}
	return result, nil
}
