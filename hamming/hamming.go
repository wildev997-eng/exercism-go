package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0

	if len(a) == len(b) {

		for i := 0; i < len(a); i++ {

			if a[i] != b[i] {
				count += 1
			}

		}

	} else {
		return 0, errors.New("strings must be the same length")
	}

	return count, nil
}
