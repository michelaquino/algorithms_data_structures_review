package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/subsets/

func main() {
	// nums := []int{1, 2, 3}
	nums := []int{9, 0, 3, 5, 7}
	answer := subsets(nums)
	fmt.Println("answer: ", answer)
}

func subsets(nums []int) [][]int {
	var buildSubset func(nums []int, partial []int, position int)

	answer := [][]int{}

	buildSubset = func(nums []int, partial []int, position int) {
		if position > len(nums)-1 {
			partialSorted := make([]int, len(partial))
			copy(partialSorted, partial)

			// only need this sort to pass on leetcode
			sort.Slice(partialSorted, func(i, j int) bool {
				return partialSorted[i] < partialSorted[j]
			})

			answer = append(answer, partialSorted)
			return
		}

		possibilities := []bool{false, true}

		for _, p := range possibilities {
			if p == true {
				buildSubset(nums, append(partial, nums[position]), position+1)
				continue
			}

			buildSubset(nums, partial, position+1)
		}
	}

	buildSubset(nums, []int{}, 0)
	return answer
}
