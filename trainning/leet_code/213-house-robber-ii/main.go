package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/house-robber-ii
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{2, 3, 2},
			expectdOutput: 3,
		},
		{
			nums:          []int{1, 2, 3, 1},
			expectdOutput: 4,
		},
		{
			nums:          []int{1, 2, 3},
			expectdOutput: 3,
		},
		// {
		// 	nums:          []int{200, 3, 140, 20, 10},
		// 	expectdOutput: 340,
		// },
		{
			nums:          []int{2, 1, 1, 2},
			expectdOutput: 3,
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

	max1 := rob_(nums, 0, len(nums)-2)
	max2 := rob_(nums, 1, len(nums)-1)

	return max(max1, max2)
}

func rob_(nums []int, start, end int) int {
	t1 := 0
	t2 := 0

	for i := start; i <= end; i++ {
		current := nums[i]

		temp := t1
		t1 = max(t1, current+t2)
		t2 = temp
	}

	return t1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
