package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/combination-sum-ii
func main() {
	testCases := []struct {
		candidates    []int
		target        int
		expectdOutput [][]int
	}{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expectdOutput: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expectdOutput: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := combinationSum2(t.candidates, t.target)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v\ngot %v", t.expectdOutput, output))
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	answer := [][]int{}
	var backtracking func(candidatesCounter [][]int, target int, current int, actualSet []int, actualSum int)

	backtracking = func(candidatesCounter [][]int, target int, current int, actualSet []int, actualSum int) {
		if actualSum > target {
			return
		}

		if actualSum == target {
			actualSetCopy := make([]int, len(actualSet))
			copy(actualSetCopy, actualSet)
			answer = append(answer, actualSetCopy)
		}

		for i := current; i < len(candidatesCounter); i++ {
			candidate, frequency := candidatesCounter[i][0], candidatesCounter[i][1]
			if frequency <= 0 {
				continue
			}

			actualSet = append(actualSet, candidate)
			candidatesCounter[i][1]--

			backtracking(candidatesCounter, target, i, actualSet, actualSum+candidate)

			candidatesCounter[i][1]++
			actualSet = actualSet[:len(actualSet)-1]
		}
	}

	candidatesCountMap := map[int]int{}
	for _, n := range candidates {
		candidatesCountMap[n]++
	}

	counter := [][]int{}
	for candidate, frequency := range candidatesCountMap {
		counter = append(counter, []int{candidate, frequency})
	}

	backtracking(counter, target, 0, []int{}, 0)
	return answer
}
