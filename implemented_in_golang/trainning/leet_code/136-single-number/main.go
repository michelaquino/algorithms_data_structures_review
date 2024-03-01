package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/single-number
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{2, 2, 1},
			expectdOutput: 1,
		},
		{
			nums:          []int{4, 1, 2, 1, 2},
			expectdOutput: 4,
		},
		{
			nums:          []int{1},
			expectdOutput: 1,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := singleNumber(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result = result ^ n

	}

	return result
}
