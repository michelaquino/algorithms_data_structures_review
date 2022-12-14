package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/

func main() {
	testCases := []struct {
		nums          []int
		expectdOutput []int
	}{
		{
			nums:          []int{-4, -1, 0, 3, 10},
			expectdOutput: []int{0, 1, 9, 16, 100},
		},
		{
			nums:          []int{-7, -3, 2, 3, 11},
			expectdOutput: []int{4, 9, 9, 49, 121},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := sortedSquaresTwoPointers(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput, output))
		}
	}

}

func sortedSquaresTwoPointers(nums []int) []int {
	awnser := []int{}
	left := 0
	right := len(nums) - 1

	for left <= right {
		number := 0
		if int(math.Abs(float64(nums[left]))) > nums[right] {
			number = nums[left]
			left++
		} else {
			number = nums[right]
			right--
		}

		awnser = append(awnser, number*number)
	}

	sort.Slice(awnser, func(i, j int) bool {
		return awnser[i] < awnser[j]
	})

	return awnser
}

func sortedSquaresSort(nums []int) []int {
	awnser := []int{}
	for _, n := range nums {
		awnser = append(awnser, n*n)
	}

	sort.Slice(awnser, func(i, j int) bool {
		return awnser[i] < awnser[j]
	})

	return awnser
}
