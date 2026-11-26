package sublist

// Relation type is defined in relations.go file.

import "slices"

func contains(larger, smaller []int) bool {
	if len(smaller) > len(larger) {
		return false
	}

	for i := 0; i+len(smaller) <= len(larger); i++ {
		window := larger[i : i+len(smaller)]

		if slices.Equal(window, smaller) {
			return true
		}
	}

	return false
}

func Sublist(l1, l2 []int) Relation {
	switch {
	case slices.Equal(l1, l2):
		return RelationEqual
	case contains(l2, l1):
		return RelationSublist
	case contains(l1, l2):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}
