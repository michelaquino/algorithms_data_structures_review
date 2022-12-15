package main

import (
	"fmt"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
var (
	numberLetterMap = map[string][]string{
		"0": {" "},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func main() {
	letterCombinations("234")
}

func letterCombinations(digits string) {
	resultList := bfs(digits, "", 0)
	fmt.Println(resultList)
}

func bfs(digits string, parcialResult string, level int) []string {
	if level == len(digits) {
		return []string{parcialResult}
	}

	digit := string(digits[level])
	letters := numberLetterMap[digit]

	resultList := []string{}
	for _, letter := range letters {
		result := bfs(digits, parcialResult+letter, level+1)
		resultList = append(resultList, result...)
	}

	return resultList
}
