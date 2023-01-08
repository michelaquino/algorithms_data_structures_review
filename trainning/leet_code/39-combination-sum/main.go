package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/combination-sum
func main() {
	testCases := []struct {
		candidates    []int
		target        int
		expectdOutput [][]int
	}{
		{
			candidates:    []int{2, 3, 6, 7},
			target:        7,
			expectdOutput: [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates:    []int{2, 3, 5},
			target:        8,
			expectdOutput: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates:    []int{2},
			target:        1,
			expectdOutput: [][]int{},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := combinationSum(t.candidates, t.target)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v\ngot %v", t.expectdOutput, output))
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	answer := [][]int{}
	var backtracking func(candidates []int, target int, start int, actualSet []int, actualSum int)
	backtracking = func(candidates []int, target int, start int, actualSet []int, actualSum int) {
		if actualSum > target {
			return
		}

		if actualSum == target {
			actualSetCopy := make([]int, len(actualSet))
			copy(actualSetCopy, actualSet)
			answer = append(answer, actualSetCopy)
		}

		for i := start; i < len(candidates); i++ {
			backtracking(candidates, target, i, append(actualSet, candidates[i]), actualSum+candidates[i])
		}
	}

	backtracking(candidates, target, 0, []int{}, 0)
	return answer
}

// func permuteUnique(nums []int) [][]int {
// 	answer := [][]int{}

// 	var backtracking func(nums, subset []int, numCounterMap map[int]int)
// 	backtracking = func(nums, subset []int, numCounterMap map[int]int) {
// 		if len(subset) == len(nums) {
// 			temp := make([]int, len(subset))
// 			copy(temp, subset)
// 			answer = append(answer, temp)
// 			return
// 		}

// 		for v, _ := range numCounterMap {
// 			if numCounterMap[v] > 0 {
// 				subset = append(subset, v)
// 				numCounterMap[v]--

// 				backtracking(nums, subset, numCounterMap)

// 				subset = subset[:len(subset)-1]
// 				numCounterMap[v]++
// 			}
// 		}
// 	}

// 	numCounterMap := map[int]int{}
// 	for _, n := range nums {
// 		numCounterMap[n]++
// 	}

// 	backtracking(nums, []int{}, numCounterMap)
// 	return answer
// }
