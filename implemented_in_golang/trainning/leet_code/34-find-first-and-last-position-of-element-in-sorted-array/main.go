package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func main() {
	testCases := []struct {
		nums          []int
		target        int
		expectdOutput []int
	}{
		{
			nums:          []int{5, 7, 7, 8, 8, 10},
			target:        8,
			expectdOutput: []int{3, 4},
		},
		{
			nums:          []int{5, 7, 7, 8, 8, 10},
			target:        6,
			expectdOutput: []int{-1, -1},
		},
		{
			nums:          []int{},
			target:        0,
			expectdOutput: []int{-1, -1},
		},
		{
			nums:          []int{3, 3, 3},
			target:        3,
			expectdOutput: []int{0, 2},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := searchRange(t.nums, t.target)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func searchRange(nums []int, target int) []int {
	a, b := divideAndConquer(nums, target, 0, len(nums)-1)
	return []int{a, b}
}

func divideAndConquer(nums []int, target, start, end int) (int, int) {
	if start > end {
		return -1, -1
	}

	if start == end {
		if nums[start] == target {
			return start, end
		}

		return -1, -1
	}

	middle := (end + start) / 2

	// left
	a, b := divideAndConquer(nums, target, start, middle)

	// right
	c, d := divideAndConquer(nums, target, middle+1, end)
	if a == -1 && b == -1 {
		return c, d
	}

	if c == -1 && d == -1 {
		return a, b
	}

	return min(a, b), max(c, d)
}

func min(a, b int) int {
	if a != -1 && a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a != -1 && a > b {
		return a
	}

	return b
}
