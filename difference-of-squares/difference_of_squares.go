package differenceofsquares

func SquareOfSum(n int) int {
	num := 1
	for i := 2; i <= n; i++ {
		num += i
	}
	return num * num
}

func SumOfSquares(n int) int {
	num := 0
	for i := 1; i <= n; i++ {
		num += i * i
	}
	return num
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
