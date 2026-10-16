package listops

import "slices"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, value := range s {
		acc = fn(acc, value)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for _, value := range slices.Backward(s) {
		acc = fn(value, acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}

	for _, n := range s {
		if fn(n) == true {
			result = append(result, n)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := IntList{}
	for _, value := range s {
		result = append(result, fn(value))
	}
	return result
}

func (s IntList) Reverse() IntList {
	slices.Reverse(s)
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, group := range lists {
		s = append(s, group...)
	}
	return s
}
