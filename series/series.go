package series

func All(n int, s string) []string {
	if n <= 0 {
		return []string{}
	}
	container := []string{}

	for i := 0; i+n <= len(s); i++ {
		container = append(container, s[i:i+n])
	}
	return container
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
