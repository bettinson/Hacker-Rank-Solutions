package main

import "testing"

func TestArrayCalculation(t *testing.T) {
	n3 := 3
	arr3 := [][]int{{2, 3, 4}, {3, 8, 6}, {2, 4, 1}}

	diff3 := diagonalDifference(arr3, n3)
	if diff3 != 11-14 {
		t.Errorf("%d does not equal actual value, %d", diff3, -3)
	}

	n4 := 4
	arr4 := [][]int{{4, 5, 2, 5}, {3, 7, 5, 9}, {4, 7, 3, 5}, {9, 1, 2, 4}}

	diff4 := diagonalDifference(arr4, n4)
	if diff4 != (4+7+3+4)-(5+5+7+9) {
		t.Errorf("%d does not equal actual value, %d", diff4, -8)
	}
}
