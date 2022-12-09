package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/array-partition/
func main() {
	// nums := []int{1, 4, 3, 2}
	nums := []int{6, 2, 6, 5, 1, 2}

	fmt.Println(arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
