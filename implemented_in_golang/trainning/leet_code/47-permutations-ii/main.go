package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/permutations-ii
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput [][]int
	}{
		{
			nums:          []int{1, 2, 3},
			expectdOutput: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{1, 1, 2},
			expectdOutput: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := permuteUnique(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v\ngot %v", t.expectdOutput, output))
		}
	}
}

func permuteUnique(nums []int) [][]int {
	answer := [][]int{}

	var backtracking func(nums, subset []int, numCounterMap map[int]int)
	backtracking = func(nums, subset []int, numCounterMap map[int]int) {
		if len(subset) == len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			answer = append(answer, temp)
			return
		}

		for v, _ := range numCounterMap {
			if numCounterMap[v] > 0 {
				subset = append(subset, v)
				numCounterMap[v]--

				backtracking(nums, subset, numCounterMap)

				subset = subset[:len(subset)-1]
				numCounterMap[v]++
			}
		}
	}

	numCounterMap := map[int]int{}
	for _, n := range nums {
		numCounterMap[n]++
	}

	backtracking(nums, []int{}, numCounterMap)
	return answer
}
