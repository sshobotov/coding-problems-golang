package day33

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	result := Solve([]int{2, 1, 5, 7, 2, 0, 5})
	if !reflect.DeepEqual(result, []float32{2, 1.5, 2, 3.5, 2, 2, 2}) {
		t.Errorf("Nooooo %v", result)
	}
}

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	heap.Push(10)
	heap.Push(1)
	heap.Push(15)

	if heap.Get() != 1 {
		t.Errorf("Failed to get correct min heap root, got %v", heap.Get())
	}

	heap.Pop()

	if heap.Get() != 10 {
		t.Errorf("Failed to get correct min heap root after pop, got %v", heap.Get())
	}
}

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()

	heap.Push(10)
	heap.Push(1)
	heap.Push(15)

	if heap.Get() != 15 {
		t.Errorf("Failed to get correct min heap root, got %v", heap.Get())
	}

	heap.Pop()

	if heap.Get() != 10 {
		t.Errorf("Failed to get correct min heap root after pop, got %v", heap.Get())
	}
}
