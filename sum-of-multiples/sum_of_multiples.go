package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	sums := 0
	seen := make(map[int]bool)
	for _, index := range divisors {
		if index <= 0 {
			continue
		}
		for i := 1; i*index < limit; i++ {
			multiple := index * i
			if !seen[multiple] {
				sums += multiple
				seen[multiple] = true
			}
		}
	}
	return sums
}
