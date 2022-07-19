package main

import (
	"testing"
)

func TestSh(t *testing.T) {

	arr := []int{1, 2, 3, 4}
	original := make([]int, len(arr))
	copy(original, arr)
	t.Run("test shaffle", func(t *testing.T) {
		shafflee(arr)
		if len(arr) != len(original) {
			t.Errorf("arr: %v, and original: %v are not the same length", arr, original)
		}
		if isEqual(arr, original) {
			t.Errorf("arr: %v, and original: %v are the same", arr, original)
		}
	})
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
