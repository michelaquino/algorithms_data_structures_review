package main

import "fmt"

// https://leetcode.com/problems/binary-search

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	// target := 2

	fmt.Println("search: ", search(nums, target))
}

func search(nums []int, target int) int {
	return bynarySearch(nums, 0, len(nums)-1, target)
}

func bynarySearch(nums []int, startIndex, endIndex, target int) int {
	if startIndex < 0 || endIndex > len(nums)-1 || startIndex > endIndex {
		return -1
	}

	middle := (endIndex + startIndex) / 2
	if nums[middle] == target {
		return middle
	}

	if target > nums[middle] {
		return bynarySearch(nums, middle+1, endIndex, target)
	}

	return bynarySearch(nums, 0, middle-1, target)
}
