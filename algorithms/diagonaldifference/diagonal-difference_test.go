package main

import "testing"

func TestArrayCalculation(t *testing.T) {
	n3 := 3
	arr3 := [][]int{{2, 3, 4}, {3, 8, 6}, {2, 4, 1}}

	diff := diagonalDifference(arr3, n3)
	if diff != 11-14 {
		t.Errorf("%d does not equal actual value, %d", diff, -3)
	}
}
