package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{3, 4, 5, 1, 2},
			expectdOutput: 1,
		},
		{
			nums:          []int{1, 2, 3, 4, 5},
			expectdOutput: 1,
		},
		{
			nums:          []int{5, 1, 2, 3, 4},
			expectdOutput: 1,
		},
		{
			nums:          []int{2, 3, 4, 5, 1},
			expectdOutput: 1,
		},
		{
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			expectdOutput: 0,
		},
		{
			nums:          []int{11, 13, 15, 17},
			expectdOutput: 11,
		},
		{
			nums:          []int{5},
			expectdOutput: 5,
		},
	}

	for _, t := range testCases {
		// o := findRotationIndex(t.nums)
		// fmt.Printf("output: '%v'\n", o)
		// continue

		fmt.Println("=========================================")
		output := findMin(t.nums)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	lastIndex := len(nums) - 1
	left := 0
	right := lastIndex

	if nums[0] < nums[lastIndex] {
		return nums[0]
	}

	for left <= right {
		pivot := (left + right) / 2
		if pivot+1 <= lastIndex && nums[pivot] > nums[pivot+1] {
			return nums[pivot+1]
		}

		if nums[pivot] < nums[left] {
			right = pivot - 1
			continue
		}

		left = pivot + 1

	}

	return -1
}
