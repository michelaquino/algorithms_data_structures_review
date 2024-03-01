package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/permutations
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput [][]int
	}{
		{
			nums:          []int{1, 2, 3},
			expectdOutput: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums:          []int{0, 1},
			expectdOutput: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums:          []int{1},
			expectdOutput: [][]int{{1}},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := permute(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func permute(nums []int) [][]int {
	answer := [][]int{}
	maxNums := len(nums)

	var backtracking func(nums, subset []int)
	backtracking = func(nums, subset []int) {

		if len(subset) == maxNums {
			answer = append(answer, subset)
			return
		}

		for i, n := range nums {
			numsCopied := make([]int, len(nums))
			copy(numsCopied, nums)

			numsWithoutCurrent := append(numsCopied[:i], numsCopied[i+1:]...)

			backtracking(numsWithoutCurrent, append(subset, n))
		}
	}

	backtracking(nums, []int{})
	return answer
}
