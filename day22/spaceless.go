package day22

/*
Given a dictionary of words and a string made up of those words (no spaces), return the original sentence in a list. If there is more than one possible reconstruction, return any of them. If there is no possible reconstruction, then return null.

For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox", you should return ['the', 'quick', 'brown', 'fox'].

Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond', and the string "bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].
*/

import "strings"

// Solve finds possible original sentence
func Solve(dict []string, s string) []string {
	return solveRecursive(dict, s)
}

func solveRecursive(dict []string, s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	for _, word := range dict {
		if strings.HasPrefix(s, word) {
			substr := s[len(word):]
			if nested := solveRecursive(dict, substr); nested != nil {
				return append(append([]string{}, word), nested...)
			}
		}
	}
	return nil
}
