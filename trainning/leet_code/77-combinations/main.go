package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/combinations
func main() {
	testCases := []struct {
		n             int
		k             int
		expectdOutput [][]int
	}{
		{
			n:             4,
			k:             2,
			expectdOutput: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:             1,
			k:             1,
			expectdOutput: [][]int{{1}},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := combine(t.n, t.k)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func combine(n int, k int) [][]int {
	answer := [][]int{}

	var backtracking func(subset []int, start, end, k int)
	backtracking = func(subset []int, start, end, k int) {
		if len(subset) == k {
			answer = append(answer, subset)
			return
		}

		for i := start; i <= end; i++ {
			setCopied := make([]int, len(subset))
			copy(setCopied, subset)

			backtracking(append(setCopied, i), i+1, end, k)
		}
	}

	backtracking([]int{}, 1, n, k)
	return answer
}
