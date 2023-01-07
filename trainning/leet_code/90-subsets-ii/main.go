package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/subsets-ii
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput [][]int
	}{
		{
			nums: []int{1, 2, 2},
			expectdOutput: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		// {
		// 	nums: []int{0},
		// 	expectdOutput: [][]int{
		// 		{},
		// 		{0},
		// 	},
		// },
		// {
		// 	nums: []int{4, 4, 4, 1, 4},
		// 	expectdOutput: [][]int{
		// 		{},
		// 		{1},
		// 		{1, 4},
		// 		{1, 4, 4},
		// 		{1, 4, 4, 4},
		// 		{1, 4, 4, 4, 4},
		// 		{4},
		// 		{4, 4},
		// 		{4, 4, 4},
		// 		{4, 4, 4, 4},
		// 	},
		// },
		{
			nums: []int{2, 1, 2, 1, 3},
			expectdOutput: [][]int{
				{},
				{1},
				{1, 1},
				{1, 1, 2},
				{1, 1, 2, 2},
				{1, 1, 2, 2, 3},
				{1, 1, 2, 3},
				{1, 1, 3},
				{1, 2},
				{1, 2, 2},
				{1, 2, 2, 3},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 2},
				{2, 2, 3},
				{2, 3},
				{3},
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := subsetsWithDup(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v, \ngot %v", t.expectdOutput, output))
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	answer := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println("nums: ", nums)

	var backtracking func(nums []int, position int, currentSubset []int)
	backtracking = func(nums []int, index int, currentSubset []int) {
		answer = append(answer, currentSubset)

		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}

			currentSubsetCopy := make([]int, len(currentSubset))
			copy(currentSubsetCopy, currentSubset)
			currentSubsetCopy = append(currentSubsetCopy, nums[i])

			backtracking(nums, i+1, currentSubsetCopy)
		}
	}

	backtracking(nums, 0, []int{})
	return answer
}

func subsetsWithDup2(nums []int) [][]int {
	setFound := map[string]bool{}
	answer := [][]int{}

	var backtracking func(nums []int, position int, partialAwnser []int)
	backtracking = func(nums []int, position int, partialAwnser []int) {
		if position > len(nums)-1 {
			partialAwnserSorted := make([]int, len(partialAwnser))
			copy(partialAwnserSorted, partialAwnser)
			// only need this sort to pass on leetcode
			sort.Slice(partialAwnserSorted, func(i, j int) bool {
				return partialAwnserSorted[i] < partialAwnserSorted[j]
			})

			key := getKey(partialAwnserSorted)
			if setFound[key] {
				return
			}

			setFound[key] = true

			answer = append(answer, partialAwnserSorted)
			return
		}

		possibilities := []bool{true, false}
		for _, possibility := range possibilities {
			num := nums[position]

			if possibility {
				backtracking(nums, position+1, append(partialAwnser, num))
				continue
			}

			backtracking(nums, position+1, partialAwnser)
		}
	}

	backtracking(nums, 0, []int{})
	return answer
}

func getKey(set []int) string {
	key := strings.Builder{}
	for _, n := range set {
		key.WriteString(strconv.Itoa(n))
		key.WriteString("#")
	}

	return key.String()
}
