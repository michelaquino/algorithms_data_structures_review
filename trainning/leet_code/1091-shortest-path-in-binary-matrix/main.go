package main

import (
	"fmt"
	"reflect"
)

var (
	notProcessed              = 0
	pieceOfIslandNotProcessed = 1
	alreadyProcessed          = -1
	islandAlreadyFound        = 2
)

// https://leetcode.com/problems/shortest-path-in-binary-matrix
func main() {
	testCases := []struct {
		grid          [][]int
		expectdOutput int
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expectdOutput: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			expectdOutput: 4,
		},
		{
			grid: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			expectdOutput: -1,
		},
		{
			grid: [][]int{
				{0},
			},
			expectdOutput: 1,
		},
	}

	for _, t := range testCases {
		// ans := findPath(t.grid, 0, 3)
		// fmt.Println(ans)
		// continue

		fmt.Println("=========================================")
		output := shortestPathBinaryMatrix(t.grid)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}

	grid[0][0] = 1
	queue := [][]int{{0, 0}}
	for len(queue) > 0 {
		row, column := queue[0][0], queue[0][1]
		distance := grid[row][column]
		queue = queue[1:]

		if row == n-1 && column == n-1 {
			return distance
		}

		directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for _, direction := range directions {
			newRow := row + direction[0]
			newColumn := column + direction[1]

			// out of limits
			if newRow < 0 || newColumn < 0 || newRow > n-1 || newColumn > n-1 {
				continue
			}

			// already processed or block
			if grid[newRow][newColumn] != 0 {
				continue
			}

			grid[newRow][newColumn] = distance + 1
			queue = append(queue, []int{newRow, newColumn})
		}
	}

	return -1
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}

		fmt.Println()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
