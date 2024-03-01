package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string
func main() {
	testCases := []struct {
		nums          []int
		k             int
		expectdOutput int
	}{
		{
			nums:          []int{10, 5, 2, 6},
			k:             100,
			expectdOutput: 8,
		},
		{
			nums:          []int{1, 2, 3},
			k:             0,
			expectdOutput: 0,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := numSubarrayProductLessThanK(t.nums, t.k)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	awnser := 0
	start := 0
	end := 0
	prod := 1

	for start <= len(nums)-1 {
		prod *= nums[end]

		if prod < k {
			awnser++
			end++
		}

		if end < len(nums) {
			continue
		}

		start++
		end = start
		prod = 1
	}

	return awnser
}
