package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var total float64 = 0
	lengthNumber := strconv.Itoa(n)
	for i := 0; i <= len(lengthNumber)-1; i++ {
		converting := int(lengthNumber[i]) - '0'
		total += math.Pow(float64(converting), float64(len(lengthNumber)))
	}
	return int(total) == n
}
