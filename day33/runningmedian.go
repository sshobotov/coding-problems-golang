package day33

/*
Compute the running median of a sequence of numbers. That is, given a stream of numbers, print out the median of the list so far on each new element.

Recall that the median of an even-numbered list is the average of the two middle numbers.
*/

// Solve returns median for running numbers
func Solve(stream []int) []float32 {
	max := NewMaxHeap()
	min := NewMinHeap()

	var result []float32
	var median float32

	for _, v := range stream {
		if max.Size() == min.Size() {
			if float32(v) > median {
				min.Push(v)
				median = float32(min.Get())
			} else {
				max.Push(v)
				median = float32(max.Get())
			}
		} else {
			if max.Size() > min.Size() {
				if float32(v) < median {
					min.Push(max.Pop())
					max.Push(v)
				} else {
					min.Push(v)
				}
			} else {
				if float32(v) > median {
					max.Push(min.Pop())
					min.Push(v)
				} else {
					max.Push(v)
				}
			}
			median = float32(min.Get()+max.Get()) / 2
		}
		result = append(result, median)
	}

	return result
}
