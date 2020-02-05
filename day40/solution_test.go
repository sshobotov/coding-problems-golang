package day40

import "testing"

func TestSolution(t *testing.T) {
	var solution int

	solution = Solve([]int{6, 1, 3, 3, 3, 6, 6})
	if solution != 1 {
		t.Errorf("Expected 1 but got %d", solution)
	}

	solution = Solve([]int{13, 19, 13, 13})
	if solution != 19 {
		t.Errorf("Expected 19 but got %d", solution)
	}
}
