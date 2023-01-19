package main

import (
	"fmt"
	"reflect"
)

/*
	aab
	ab
	b
	aab

*/

// https://leetcode.com/problems/palindrome-partitioning/
func main() {
	// s := "aab"
	// fmt.Println(isPalindrome(s, 0, 1))
	// fmt.Println(s[0:2])
	// return

	testCases := []struct {
		s             string
		expectdOutput [][]string
	}{
		{
			s: "aab",
			expectdOutput: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			s: "a",
			expectdOutput: [][]string{
				{"a"},
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := partition(t.s)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func partition(s string) [][]string {
	awnser := [][]string{}

	var backtracking func(s string, currentList []string, start int)
	backtracking = func(s string, currentList []string, start int) {
		if start >= len(s) {
			currentListCopy := make([]string, len(currentList))
			copy(currentListCopy, currentList)
			awnser = append(awnser, currentListCopy)
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				currentList = append(currentList, s[start:end+1])
				backtracking(s, currentList, end+1)
				currentList = currentList[:len(currentList)-1]
			}
		}
	}

	backtracking(s, []string{}, 0)
	return awnser
}

func isPalindrome(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
