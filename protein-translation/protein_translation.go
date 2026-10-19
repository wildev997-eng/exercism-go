package proteintranslation

import (
	"errors"
)

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid codon")
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	listRNA := []string{}

	if rna == "" {
		return listRNA, nil
	}

	if len(rna) == 4 {
		return listRNA, ErrInvalidBase
	}

	for i := 0; i+3 <= len(rna); i += 3 {
		codon, exist := codons[rna[i:i+3]]

		if codon == "STOP" {
			break
		}

		if !exist {
			return listRNA, ErrInvalidBase
		}

		listRNA = append(listRNA, codon)
	}
	return listRNA, nil
}

func FromCodon(codon string) (string, error) {
	amino, exists := codons[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if amino == "STOP" {
		return "", ErrStop
	}
	return amino, nil
}
