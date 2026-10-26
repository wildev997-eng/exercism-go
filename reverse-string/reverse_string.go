package reversestring

func Reverse(input string) string {
	runes := []rune(input)
	end := len(runes) - 1
	reverse := ""
	for i := end; i >= 0; i-- {
		reverse += string(runes[i])
	}
	return reverse
}
