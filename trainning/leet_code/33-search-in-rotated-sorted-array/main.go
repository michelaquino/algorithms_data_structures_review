package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func main() {
	testCases := []struct {
		nums          []int
		target        int
		expectdOutput int
	}{
		{
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			target:        0,
			expectdOutput: 4,
		},
		{
			nums:          []int{4, 5, 6, 7, 0, 1, 2},
			target:        3,
			expectdOutput: -1,
		},
		{
			nums:          []int{1},
			target:        0,
			expectdOutput: -1,
		},
		{
			nums:          []int{3, 1},
			target:        1,
			expectdOutput: 1,
		},
		{
			nums:          []int{3, 1},
			target:        3,
			expectdOutput: 0,
		},
		{
			nums:          []int{1, 3},
			target:        3,
			expectdOutput: 1,
		},
		{
			nums:          []int{5, 1, 3},
			target:        3,
			expectdOutput: 2,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := search(t.nums, t.target)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func search(nums []int, target int) int {
	rotationIndex := findRotationIndex(nums)
	fmt.Println("rotationIndex: ", rotationIndex)

	if nums[rotationIndex] == target {
		return rotationIndex
	}

	if rotationIndex == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	if target < nums[0] {
		return binarySearch(nums, target, rotationIndex, len(nums)-1)
	}

	return binarySearch(nums, target, 0, rotationIndex)
}

func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		pivot := (left + right) / 2

		if nums[pivot] == target {
			return pivot
		}

		if nums[pivot] > target {
			right = pivot - 1
			continue
		}

		left = pivot + 1
	}

	return -1
}

func findRotationIndex(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		pivot := (left + right) / 2
		if len(nums)-1 >= pivot+1 && nums[pivot] > nums[pivot+1] {
			return pivot + 1
		}

		if nums[pivot] < nums[left] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return 0
}

// func divideAndConquer(nums []int, target, start, end int) int {
// 	if start > end {
// 		return -1
// 	}

// 	if start == end {
// 		if nums[start] == target {
// 			return start
// 		}

// 		return -1
// 	}

// 	middle := (end + start) / 2
// 	left := divideAndConquer(nums, target, start, middle)
// 	right := divideAndConquer(nums, target, middle+1, end)

// 	return max(left, right)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }
