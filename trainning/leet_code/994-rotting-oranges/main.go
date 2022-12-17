package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/rotting-oranges/
func main() {
	testCases := []struct {
		matrix        [][]int
		expectdOutput int
	}{
		{
			matrix:        [][]int{{1, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expectdOutput: -1,
		},
		{
			matrix:        [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expectdOutput: 4,
		},
		{
			matrix:        [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expectdOutput: -1,
		},
		{
			matrix:        [][]int{{0, 1}},
			expectdOutput: -1,
		},
		{
			matrix:        [][]int{{0, 2}},
			expectdOutput: 0,
		},
		{
			matrix:        [][]int{{0}},
			expectdOutput: 0,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := orangesRotting(t.matrix)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func printMatrix(matrix [][]int) {
	fmt.Println()
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println()
	}
}

type index struct {
	row    int
	column int
}

func orangesRotting(grid [][]int) int {
	freshOranges := 0
	queue := []index{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			state := grid[i][j]

			if state == 1 {
				freshOranges++
			} else if state == 2 {
				queue = append(queue, index{i, j})
			}
		}
	}

	queue = append(queue, index{-1, -1}) // ticker
	minutes := -1

	for len(queue) > 0 {
		toProcess := queue[0]
		queue = queue[1:]

		if toProcess.row == -1 {
			minutes++

			if len(queue) > 0 {
				queue = append(queue, index{-1, -1})
			}

			continue
		}

		directions := []index{
			{row: -1, column: 0},
			{row: 1, column: 0},
			{row: 0, column: -1},
			{row: 0, column: 1},
		}

		for _, direction := range directions {
			neighboorRow := toProcess.row + direction.row
			neighboorColumn := toProcess.column + direction.column

			if !insideLimit(grid, neighboorRow, neighboorColumn) {
				continue
			}

			if grid[neighboorRow][neighboorColumn] == 1 {
				grid[neighboorRow][neighboorColumn] = 2
				freshOranges--

				n := index{neighboorRow, neighboorColumn}
				queue = append(queue, n)
			}

		}
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes
}

func insideLimit(grid [][]int, i, j int) bool {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[i])-1 {
		return false
	}

	return true
}
