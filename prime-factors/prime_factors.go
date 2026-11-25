package primefactors

import "math"

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(3); i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Factors(n int64) []int64 {
	numbers := []int64{}

	for i := int64(2); n > 1; {
		if isPrime(i) {
			if n%i == 0 {
				n = n / i
				numbers = append(numbers, i)
			} else {
				i++
			}
		} else {
			i++
		}
	}

	return numbers

}
