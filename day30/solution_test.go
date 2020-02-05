package day30

import "testing"

func TestSolution(t *testing.T) {
	var solved uint

	solved = Solve([]uint{2, 1, 2})
	if solved != 1 {
		t.Errorf("Expected to get 1 got %d", solved)
	}

	solved = Solve([]uint{3, 0, 1, 3, 0, 5})
	if solved != 8 {
		t.Errorf("Expected to get 8 got %d", solved)
	}

	solved = Solve([]uint{3, 0, 1, 3, 0, 1, 0})
	if solved != 6 {
		t.Errorf("Expected to get 6 got %d", solved)
	}
}
