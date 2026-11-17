package nucleotidecount

import (
	"errors"
	"strings"
)

type DNA string
type Histogram map[rune]int

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		if !strings.ContainsRune("CGAT", nucleotide) {
			return nil, errors.New("invalid nucleotide")
		}
		h[nucleotide]++
	}
	return h, nil
}
