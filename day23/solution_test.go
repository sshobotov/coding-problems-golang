package day23

import "testing"

func TestSolution(t *testing.T) {
	solution1 := Solve([][]bool{
		[]bool{false, false, false, false},
		[]bool{true, true, false, true},
		[]bool{false, false, false, false},
		[]bool{false, false, false, false}}, &Point{3, 0}, &Point{0, 0})

	if solution1 != 7 {
		t.Errorf("Valid 7 steps should be found but received %d", solution1)
	}

	solution2 := Solve([][]bool{
		[]bool{false, false, false, true},
		[]bool{false, true, false, false},
		[]bool{false, false, true, false},
		[]bool{true, false, false, false}}, &Point{0, 2}, &Point{2, 0})

	if solution2 != 4 {
		t.Errorf("Valid 4 steps should be found but received %d", solution2)
	}
}
