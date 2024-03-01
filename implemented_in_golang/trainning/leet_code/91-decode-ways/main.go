package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/decode-ways
func main() {
	testCases := []struct {
		s             string
		expectdOutput int
	}{
		{
			s:             "11106",
			expectdOutput: 2,
		},
		{
			s:             "12",
			expectdOutput: 2,
		},
		{
			s:             "226",
			expectdOutput: 3,
		},
		{
			s:             "06",
			expectdOutput: 0,
		},
		{
			s:             "111111111111111111111111111111111111111111111",
			expectdOutput: 0,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := numDecodings(t.s)
		fmt.Println("output: ", output)

		// if !reflect.DeepEqual(output, t.expectdOutput) {
		// 	panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		// }
	}
}

func numDecodings(s string) int {
	var dynamicProgramming func(str string, i int, memo []int) int

	dynamicProgramming = func(str string, i int, memo []int) int {
		if i == len(str) {
			return 1
		}

		if memo[i] != 0 {
			return memo[i]
		}

		answer := 0
		firstChar := str[i]
		if firstChar != '0' {
			answer += dynamicProgramming(str, i+1, memo)
		}

		if i+1 < len(s) {
			secondChar, _ := strconv.Atoi(string(str[i+1]))
			if firstChar == '1' || (firstChar == '2' && secondChar <= 6) {
				answer += dynamicProgramming(str, i+2, memo)
			}
		}

		memo[i] = answer
		return answer
	}

	memo := make([]int, len(s))
	return dynamicProgramming(s, 0, memo)
}
