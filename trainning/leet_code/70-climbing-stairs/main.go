package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/climbing-stairs
func main() {
	testCases := []struct {
		n             int
		expectdOutput int
	}{
		{
			n:             2,
			expectdOutput: 2,
		},
		{
			n:             3,
			expectdOutput: 3,
		},
		{
			n:             4,
			expectdOutput: 5,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := climbStairs(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func climbStairs(n int) int {
	memo := make([]int, n+1)

	var recursive func(n int, memo []int) int
	recursive = func(n int, memo []int) int {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		}

		if memo[n] > 0 {
			return memo[n]
		}

		result := recursive(n-1, memo) + recursive(n-2, memo)
		memo[n] = result
		return result
	}

	return recursive(n, memo)
}
