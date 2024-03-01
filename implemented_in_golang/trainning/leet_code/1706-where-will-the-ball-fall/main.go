package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/where-will-the-ball-fall
func main() {
	testCases := []struct {
		grid          [][]int
		expectdOutput []int
	}{
		{
			grid:          [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}},
			expectdOutput: []int{1, -1, -1, -1, -1},
		},
		{
			grid:          [][]int{{-1}},
			expectdOutput: []int{-1},
		},
		{
			grid:          [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}},
			expectdOutput: []int{0, 1, 2, 3, 4, -1},
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.grid, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := findBall(t.grid)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func findBall(grid [][]int) []int {
	columns := len(grid[0])
	awnser := make([]int, columns)

	for column := 0; column < columns; column++ {
		result := findBallOnColumn(grid, column)
		awnser[column] = result
	}

	return awnser
}

func findBallOnColumn(grid [][]int, column int) int {
	rows := len(grid)
	columns := len(grid[0])

	row := 0
	for row < rows {
		value := grid[row][column]
		switch value {
		case 1:
			// stuck on right border or get stuck in a V
			if column+1 >= columns || grid[row][column+1] == -1 {
				return -1
			}

			row = row + 1
			column = column + 1

		case -1:
			// stuck on left border or get stuck in a V
			if column-1 < 0 || grid[row][column-1] == 1 {
				return -1
			}

			row = row + 1
			column = column - 1
		}

	}

	return column
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}

		fmt.Println()
	}
}
