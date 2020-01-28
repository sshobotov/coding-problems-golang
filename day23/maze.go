package day23

// Board represents a board. Each True boolean represents a wall. Each False boolean represents a tile you can walk on
type Board [][]bool

// Point represents coordinates on the bord
type Point struct {
	Row int
	Col int
}

// Steps represents how many steps it took to reach end of the board
type Steps int

// Solve returns the minimum number of steps required to reach the end coordinate from the start. If there is no possible path, then returns -1. You can move up, left, down, and right. You cannot move through walls. You cannot wrap around the edges of the board
func Solve(board Board, start *Point, end *Point) Steps {
	m := len(board)
	if m < 1 {
		return -1
	}
	n := len(board[0])
	if n < 1 {
		return -1
	}

	memo := make(map[Point]int)
	return Steps(solveRec(board, *start, end, m, n, make(map[Point]struct{}), &memo))
}

func solveRec(board Board, current Point, target *Point, m int, n int, visited map[Point]struct{}, memo *map[Point]int) int {
	if current == *target {
		return 0
	}
	if board[current.Row][current.Col] {
		return -1
	}
	if v, ok := (*memo)[current]; ok {
		return v
	}

	candidates := []int{}
	if current.Row > 0 {
		next := Point{current.Row - 1, current.Col}
		if _, ok := visited[next]; !ok {
			candidates = append(candidates, solveRec(board, next, target, m, n, copyVisited(visited, current), memo))
		}
	}
	if current.Row < m-1 {
		next := Point{current.Row + 1, current.Col}
		if _, ok := visited[next]; !ok {
			candidates = append(candidates, solveRec(board, next, target, m, n, copyVisited(visited, current), memo))
		}
	}
	if current.Col > 0 {
		next := Point{current.Row, current.Col - 1}
		if _, ok := visited[next]; !ok {
			candidates = append(candidates, solveRec(board, next, target, m, n, copyVisited(visited, current), memo))
		}
	}
	if current.Col < n-1 {
		next := Point{current.Row, current.Col + 1}
		if _, ok := visited[next]; !ok {
			candidates = append(candidates, solveRec(board, next, target, m, n, copyVisited(visited, current), memo))
		}
	}

	minimized := minOf(candidates...)
	result := -1
	if minimized >= 0 {
		result = 1 + minimized
	}
	(*memo)[current] = result

	return result
}

// minOf looks through values for minimal value greater then -1, otherwise it will return -1
func minOf(vars ...int) int {
	if len(vars) < 1 {
		return -1
	}

	min := vars[0]
	for _, i := range vars {
		if i >= 0 && (min > i || min < 0) {
			min = i
		}
	}

	return min
}

func copyVisited(origin map[Point]struct{}, add Point) map[Point]struct{} {
	copy := make(map[Point]struct{})
	for k, v := range origin {
		copy[k] = v
	}
	copy[add] = struct{}{}

	return copy
}
