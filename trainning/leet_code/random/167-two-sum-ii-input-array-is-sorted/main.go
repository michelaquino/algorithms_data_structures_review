package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func main() {
	testCases := []struct {
		nums          []int
		target        int
		expectdOutput []int
	}{
		{
			nums:          []int{2, 7, 11, 15},
			target:        9,
			expectdOutput: []int{1, 2},
		},
		{
			nums:          []int{2, 3, 4},
			target:        6,
			expectdOutput: []int{1, 3},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := twoSum(t.nums, t.target)
		fmt.Println("output: ", t.nums)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput, output))
		}
	}
}

func twoSumIterateOverArray(numbers []int, target int) []int {
	output := []int{}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return output
}

func twoSum(numbers []int, target int) []int {
	output := []int{}

	for i := 0; i < len(numbers); i++ {
		start := i + 1
		end := len(numbers) - 1

		for end >= start {
			middle := (end + start) / 2

			if middle == i {
				break
			}

			result := numbers[i] + numbers[middle]

			if result == target {
				return []int{i + 1, middle + 1}
			} else if result > target {
				end = middle - 1
			} else {
				start = middle + 1
			}
		}
	}

	return output
}
