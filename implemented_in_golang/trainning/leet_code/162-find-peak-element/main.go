package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-peak-element
// A peak element is an element that is strictly greater than its neighbors.
// The intution is that you can find any peak
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{1, 2, 3, 1},
			expectdOutput: 2,
		},
		{
			nums:          []int{1, 2, 1, 3, 5, 6, 4},
			expectdOutput: 5,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := findPeakElement(t.nums)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func findPeakElement(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1)
}

func binarySearch(nums []int, left, right int) int {
	for left < right {
		pivot := (left + right) / 2
		if nums[pivot] < nums[pivot+1] { // ascending
			left = pivot + 1
			continue
		}

		right = pivot
	}

	return left
}
