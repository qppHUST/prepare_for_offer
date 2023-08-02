package search

import "testing"

func TestBinaruSearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	target := 2
	index := BinarySearch(data, target)
	print(index)
}
