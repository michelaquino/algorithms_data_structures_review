package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/maximum-subarray-min-product/

func main() {
	// nums := []int{3, 1, 5, 6, 4, 2} // 60
	nums := []int{1, 2, 3, 2} // 14
	// nums := []int{1, 1, 3, 2, 2, 2, 1, 5, 1, 5} // 25

	fmt.Println(maxSumMinProductTwoPointers(nums))
	fmt.Println(maxSumMinProductMonotoicStack(nums))

}

// O(N)
func maxSumMinProductMonotoicStack(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nums = append(nums, -1)
	// Build the prefix sum
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	maxSum := 0
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[0]] >= nums[i] {
			left := stack[0]
			stack = stack[1:] //pop

			psum := prefixSum[i-1]
			if len(stack) > 0 {
				psum = prefixSum[i-1] - prefixSum[stack[0]]
			}

			less := nums[left]
			maxSum = max(maxSum, less*psum)
		}

		stack = append([]int{i}, stack...)
	}

	modulo := int(1e9 + 7)
	return maxSum % modulo
}

// O(N^2)
// Time Limit Exceeded
func maxSumMinProductTwoPointers(nums []int) int {
	var awnser int = -1

	for i := 0; i < len(nums); i++ {
		minValue := math.MaxInt64
		sum := 0

		for j := i; j < len(nums); j++ {
			minValue = min(minValue, nums[j])
			sum += nums[j]

			partialValue := sum * minValue
			awnser = max(awnser, partialValue)
		}
	}

	modulo := int(1e9 + 7)
	return int(awnser % modulo)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
