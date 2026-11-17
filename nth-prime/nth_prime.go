package nthprime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Number must be greater than 0")
	}

	nth := 0
	howManyPrime := 0

	for i := 0; howManyPrime < n; i++ {
		if isPrime(i) {
			howManyPrime++
			nth = i
		}
	}

	return nth, nil
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	// Check odd divisors up to sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
