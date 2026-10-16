package cards

import "slices"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	x := slice
	if len(x)-1 >= index && index >= 0 {
		return x[index]
	} else {
		return -1
	}
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	x := slice
	if len(x)-1 >= index && index >= 0 {
		x[index] = value
		return x
	} else {
		x = append(x, value)
		return x
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	y := values
	if values == nil {
		return slice
	} else {
		y = append(y, slice...)
		return y
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	x := slice
	if len(x)-1 >= index && index >= 0 {
		x = slices.Delete(x, index, index+1)
		return x
	} else {
		return x
	}
}
