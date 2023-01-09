package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/combination-sum
func main() {
	testCases := []struct {
		n             int
		expectdOutput []string
	}{
		{
			n:             3,
			expectdOutput: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:             1,
			expectdOutput: []string{"()"},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := generateParenthesis(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v\ngot %v", t.expectdOutput, output))
		}
	}
}

func generateParenthesis(n int) []string {
	awnser := []string{}
	charCountMap := map[string]int{
		"(": n,
		")": n,
	}

	var backtracking func(last string, charCountMap map[string]int, actualResult string)
	backtracking = func(last string, charCountMap map[string]int, actualResult string) {
		if len(actualResult) == n*2 {
			awnser = append(awnser, actualResult)
			return
		}

		possibilities := []string{"(", ")"}
		for _, possibility := range possibilities {
			if charCountMap[possibility] <= 0 {
				continue
			}

			if charCountMap[")"] < charCountMap["("] {
				continue
			}

			actualResult += possibility
			charCountMap[possibility]--

			backtracking(possibility, charCountMap, actualResult)

			actualResult = actualResult[:len(actualResult)-1]
			charCountMap[possibility]++
		}

	}

	backtracking("", charCountMap, "")
	return awnser
}
