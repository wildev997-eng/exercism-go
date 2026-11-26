package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	marked := make([]bool, limit+1)
	primes := []int{}

	for number := 2; number <= limit; number++ {
		if marked[number] {
			continue
		}

		primes = append(primes, number)

		for multiple := number * 2; multiple <= limit; multiple += number {
			marked[multiple] = true
		}
	}

	return primes
}
