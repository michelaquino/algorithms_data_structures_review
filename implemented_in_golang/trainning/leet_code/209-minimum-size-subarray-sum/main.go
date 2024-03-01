package main

import (
	"fmt"
	"math"
	"reflect"
)

// https://leetcode.com/problems/minimum-size-subarray-sum
func main() {
	testCases := []struct {
		nums          []int
		target        int
		expectdOutput int
	}{
		{
			nums:          []int{2, 3, 1, 2, 4, 3},
			target:        7,
			expectdOutput: 2,
		},
		{
			nums:          []int{1, 4, 4},
			target:        4,
			expectdOutput: 1,
		},
		{
			nums:          []int{1, 1, 1, 1, 1, 1, 1, 1},
			target:        11,
			expectdOutput: 0,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := minSubArrayLen(t.target, t.nums)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func minSubArrayLen(target int, nums []int) int {
	start := 0
	end := 0

	minValue := math.MaxInt
	sum := 0

	for start < len(nums) && end < len(nums) {
		if nums[end] >= target {
			minValue = 1
			end++
			start = end
			sum = 0
			continue
		}

		sum += nums[end]
		if sum >= target {
			length := end - start + 1
			minValue = min(minValue, length)

			sum -= nums[start]
			sum -= nums[end]
			start++
			continue
		}

		end++
	}

	if minValue == math.MaxInt {
		return 0
	}

	return minValue
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
