package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/house-robber
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{1, 2, 3, 1},
			expectdOutput: 4,
		},
		{
			nums:          []int{2, 7, 9, 3, 1},
			expectdOutput: 12,
		},
		{
			nums:          []int{3},
			expectdOutput: 3,
		},
		{
			nums:          []int{},
			expectdOutput: 0,
		},
		{
			nums:          []int{2, 1},
			expectdOutput: 2,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := rob(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	memo := make([]int, len(nums)+1)

	memo[0] = 0
	memo[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		memo[i+1] = max(memo[i], memo[i-1]+nums[i])
	}

	return memo[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
