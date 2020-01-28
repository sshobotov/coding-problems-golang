package day22

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var solution []string

	solution = Solve([]string{"quick", "brown", "the", "fox"}, "thequickbrownfox")
	if !reflect.DeepEqual(solution, []string{"the", "quick", "brown", "fox"}) {
		t.Errorf("%v is not valid solution", solution)
	}

	solution = Solve([]string{"bed", "bath", "bedbath", "and", "beyond"}, "bedbathandbeyond")
	if !reflect.DeepEqual(solution, []string{"bed", "bath", "and", "beyond"}) {
		t.Errorf("%v is not valid solution", solution)
	}

	solution = Solve([]string{"bed", "bedbath", "and", "beyond"}, "bedbathandbeyond")
	if !reflect.DeepEqual(solution, []string{"bedbath", "and", "beyond"}) {
		t.Errorf("%v is not valid solution", solution)
	}
}
