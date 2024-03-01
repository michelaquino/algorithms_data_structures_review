package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		{
			nums:          []int{10, 9, 2, 5, 3, 7, 101, 18},
			expectdOutput: 4,
		},
		{
			nums:          []int{0, 1, 0, 3, 2, 3},
			expectdOutput: 4,
		},
		{
			nums:          []int{7, 7, 7, 7, 7, 7, 7},
			expectdOutput: 1,
		},
		{
			nums:          []int{4, 10, 4, 3, 8, 9},
			expectdOutput: 3,
		},
		{
			nums:          []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			expectdOutput: 6,
		},
		{
			nums:          []int{2, 6, 8, 3, 4, 5, 1},
			expectdOutput: 4,
		},
	}

	for _, t := range testCases {

		// index := indexOfSmallerNumberThatIsGreaterThanActual([]int{2, 6, 8}, 3)
		// fmt.Println("index: ", index)

		// continue
		fmt.Println("=========================================")
		output := lengthOfLIS(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func lengthOfLIS(nums []int) int {
	maxSizeSoFar := make([]int, len(nums))
	for i := 0; i < len(maxSizeSoFar); i++ {
		maxSizeSoFar[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxSizeSoFar[i] = max(maxSizeSoFar[i], maxSizeSoFar[j]+1)
			}
		}
	}

	result := -1
	for _, m := range maxSizeSoFar {
		result = max(result, m)
	}

	return result
}

func lengthOfLIS_BinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	subSequence := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		actualNumber := nums[i]
		lastNumber := subSequence[len(subSequence)-1]

		if actualNumber > lastNumber {
			subSequence = append(subSequence, actualNumber)
		} else {
			index := indexOfSmallerNumberThatIsGreaterThanActual(subSequence, actualNumber)
			if index >= 0 {
				subSequence[index] = actualNumber
			}
		}
	}

	return len(subSequence)
}

func indexOfSmallerNumberThatIsGreaterThanActual(subSequence []int, actualNumber int) int {
	start := 0
	end := len(subSequence) - 1
	possibleIndex := -1

	for start <= end {
		middle := (end + start) / 2

		if subSequence[middle] < actualNumber {
			start = middle + 1
		} else {
			possibleIndex = middle
			end = middle - 1
		}
	}

	return possibleIndex
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
