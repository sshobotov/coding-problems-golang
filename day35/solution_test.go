package day35

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	example := []rune{'G', 'B', 'R', 'R', 'B', 'R', 'G'}
	Solve(example)
	if !reflect.DeepEqual(example, []rune{'R', 'R', 'R', 'G', 'G', 'B', 'B'}) {
		t.Errorf("Unexpected result %v", example)
	}
}
