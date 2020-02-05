package day40

/*
Given an array of integers where every integer occurs three times except for one integer, which only occurs once, find and return the non-duplicated integer.

For example, given [6, 1, 3, 3, 3, 6, 6], return 1. Given [13, 19, 13, 13], return 19.

Do this in O(N) time and O(1) space.
*/

import (
	"fmt"
	"strconv"
)

// Solve solves given problem
func Solve(in []int) int {
	var acc []int
	for _, v := range in {
		bin := []rune(fmt.Sprintf("%b", v))
		binSize := len(bin)

		for i := range bin {
			num := ctoi(bin[binSize-1-i])
			if len(acc) < i+1 {
				acc = append(acc, num)
			} else {
				acc[i] += num
			}
		}
	}

	result := make([]rune, len(acc))
	for i, c := range acc {
		result[len(acc)-1-i] = itoc(c % 3)
	}

	value, _ := strconv.ParseInt(string(result), 2, 0)
	return int(value)
}

func ctoi(c rune) int {
	return int(c - '0')
}

func itoc(i int) rune {
	return rune(i + '0')
}
