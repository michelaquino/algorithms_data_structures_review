package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/triangle
func main() {
	testCases := []struct {
		triangle      [][]int
		expectdOutput int
	}{
		{
			triangle:      [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			expectdOutput: 11,
		},
		{
			triangle:      [][]int{{-10}},
			expectdOutput: -10,
		},
		{
			triangle:      [][]int{},
			expectdOutput: 0,
		},
		{
			triangle:      [][]int{{-1}, {2, 3}, {1, -1, -3}},
			expectdOutput: -1,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := minimumTotal(t.triangle)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func minimumTotal(triangle [][]int) int {
	memo := map[string]int{}
	return minPath(triangle, memo, 0, 0)
}

func minPath(triangle [][]int, memo map[string]int, row, column int) int {
	key := fmt.Sprintf("%d:%d", row, column)
	if sum, exists := memo[key]; exists {
		return sum
	}

	if row > len(triangle)-1 || column > len(triangle[row])-1 {
		return 0
	}

	actualPath := triangle[row][column]
	sum := actualPath + min(minPath(triangle, memo, row+1, column), minPath(triangle, memo, row+1, column+1))
	memo[key] = sum
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
