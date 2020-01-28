package day31

import "testing"

func TestSolution(t *testing.T) {
	var solved int

	solved = Solve("ok", "yes")
	if solved != 3 {
		t.Errorf("Expected to get 3 but %d received", solved)
	}

	solved = Solve("sitting", "kitten")
	if solved != 3 {
		t.Errorf("Expected to get 3 but %d received", solved)
	}

	solved = Solve("worker", "coworking")
	if solved != 5 {
		t.Errorf("Expected to get 5 but %d received", solved)
	}
}
