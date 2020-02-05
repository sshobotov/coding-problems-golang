package day35

/*
Given an array of strictly the characters 'R', 'G', and 'B', segregate the values of the array so that all the Rs come first, the Gs come second, and the Bs come last. You can only swap elements of the array.

Do this in linear time and in-place.

For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'], it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].
*/

// Solve solves segregation of values problem
func Solve(in []rune) {
	current := 0
	insertLeft := 0
	insertRight := len(in) - 1

	for {
		switch in[current] {
		case 'R':
			if current != insertLeft {
				in[current], in[insertLeft] = in[insertLeft], in[current]
				insertLeft++
			} else {
				current++
			}
		case 'G':
			current++
		case 'B':
			if current != insertRight {
				in[current], in[insertRight] = in[insertRight], in[current]
				insertRight--
			} else {
				current++
			}
		}
		if current > insertRight {
			break
		}
	}
}
