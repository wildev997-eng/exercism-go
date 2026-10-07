package leap

func IsLeapYear(year int) bool {
	leapYear := false
	if year%400 == 0 && year%100 == 0 {
		leapYear = true
		return leapYear
	} else if year%100 == 0 {
		return leapYear
	} else if year%4 == 0 {
		leapYear = true
		return leapYear
	} else {
		return leapYear
	}
}
