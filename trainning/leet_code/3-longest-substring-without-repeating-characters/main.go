package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func main() {
	testCases := []struct {
		str           string
		expectdOutput int
	}{
		{
			str:           "abcabcbb",
			expectdOutput: 3,
		},
		{
			str:           "bbbbb",
			expectdOutput: 1,
		},
		{
			str:           "pwwkew",
			expectdOutput: 3,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := lengthOfLongestSubstring(t.str)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func lengthOfLongestSubstring(str string) int {
	lettersCountMap := map[byte]int{}

	maxStr := 0
	first := 0
	second := 0
	for second < len(str) {
		char := str[second]
		lettersCountMap[char]++

		if repeatedChar(lettersCountMap) {
			lettersCountMap[str[first]]--
			first++
		}

		maxStr = max(maxStr, second-first+1)
		second++
	}

	return maxStr
}

func repeatedChar(lettersCountMap map[byte]int) bool {
	for _, count := range lettersCountMap {
		if count > 1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
