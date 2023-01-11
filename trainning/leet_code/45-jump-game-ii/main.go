package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/jump-game-ii
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput int
	}{
		// {
		// 	nums:          []int{2, 3, 1, 1, 4},
		// 	expectdOutput: 2,
		// },
		// {
		// 	nums:          []int{2, 3, 0, 1, 4},
		// 	expectdOutput: 2,
		// },
		// {
		// 	nums:          []int{1, 2, 3},
		// 	expectdOutput: 2,
		// },
		// {
		// 	nums:          []int{1},
		// 	expectdOutput: 0,
		// },
		// {
		// 	nums:          []int{1, 2},
		// 	expectdOutput: 1,
		// },
		{
			nums:          []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3},
			expectdOutput: 2,
		},
	}

	for _, t := range testCases {
		// ans := findPath(t.graph, 0, 3)
		// fmt.Println(ans)
		// continue

		fmt.Println("=========================================")
		output := jump(t.nums)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func jump(nums []int) int {
	farthest := 0
	currentEnd := 0
	jumps := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
