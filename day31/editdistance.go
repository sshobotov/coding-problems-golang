package day31

/*
The edit distance between two strings refers to the minimum number of character insertions, deletions, and substitutions required to change one string to the other. For example, the edit distance between “kitten” and “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”, and append a “g”.

Given two strings, compute the edit distance between them.
*/

// Solve computes the edit distance between two strings
func Solve(fst string, snd string) int {
	return countEditions([]rune(fst), []rune(snd))
}

func countEditions(fst []rune, snd []rune) int {
	if len(fst) == 0 {
		return len(snd)
	}
	if len(snd) == 0 {
		return len(fst)
	}
	if fst[0] == snd[0] {
		return countEditions(fst[1:], snd[1:])
	}

	var results []int

	results = append(results, 1+countEditions(fst[1:], snd))
	results = append(results, 1+countEditions(fst, snd[1:]))
	results = append(results, 1+countEditions(fst[1:], snd[1:]))

	min := results[0]
	for i := range results {
		if results[i] < min {
			min = results[i]
		}
	}
	return min
}
