package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/longest-common-subsequence/
// https://www.youtube.com/watch?v=ASoaQq66foQ
func main() {
	testCases := []struct {
		text1         string
		text2         string
		expectdOutput int
	}{
		{
			text1:         "abcde",
			text2:         "ace",
			expectdOutput: 3,
		},
		{
			text1:         "abc",
			text2:         "abc",
			expectdOutput: 3,
		},
		{
			text1:         "abc",
			text2:         "def",
			expectdOutput: 0,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := longestCommonSubsequence(t.text1, t.text2)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text2)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(text1)+1)
	}

	for i := 1; i < len(memo); i++ {
		for j := 1; j < len(memo[i]); j++ {
			if text2[i-1] == text1[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}

	return memo[len(text2)][len(text1)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println()
	}
}
