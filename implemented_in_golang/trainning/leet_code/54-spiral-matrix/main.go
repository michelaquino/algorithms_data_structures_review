package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/spiral-matrix
func main() {
	testCases := []struct {
		matrix        [][]int
		expectdOutput []int
	}{
		{
			matrix:        [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expectdOutput: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix:        [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expectdOutput: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix:        [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			expectdOutput: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := spiralOrder(t.matrix)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func printMatrix(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%d ", table[i][j])
		}

		fmt.Println()
	}
}

const (
	alreadyVisited = 101

	directionRight = 1
	directionDown  = 2
	directionLeft  = 3
	directionUp    = 4
)

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	row := 0
	column := 0
	direction := directionRight

	for row != -1 && column != -1 {
		result = append(result, matrix[row][column])
		matrix[row][column] = alreadyVisited
		row, column, direction = nextDirection(matrix, row, column, direction)
	}

	return result
}

func nextDirection(matrix [][]int, row, column, direction int) (nextRow int, nextColumn int, nextDirection int) {
	canGoRight := func(i, j int) bool {
		return canVisited(matrix, i, j+1)
	}

	canGoDown := func(i, j int) bool {
		return canVisited(matrix, i+1, j)
	}

	canGoLeft := func(i, j int) bool {
		return canVisited(matrix, i, j-1)
	}

	canGoUp := func(i, j int) bool {
		return canVisited(matrix, i-1, j)
	}

	switch direction {
	case directionRight:
		if canGoRight(row, column) {
			nextRow = row
			nextColumn = column + 1
			nextDirection = directionRight
			return
		} else if canGoDown(row, column) {
			nextRow = row + 1
			nextColumn = column
			nextDirection = directionDown
			return
		}
	case directionDown:
		if canGoDown(row, column) {
			nextRow = row + 1
			nextColumn = column
			nextDirection = directionDown
			return
		} else if canGoLeft(row, column) {
			nextRow = row
			nextColumn = column - 1
			nextDirection = directionLeft
			return
		}
	case directionLeft:
		if canGoLeft(row, column) {
			nextRow = row
			nextColumn = column - 1
			nextDirection = directionLeft
			return
		} else if canGoUp(row, column) {
			nextRow = row - 1
			nextColumn = column
			nextDirection = directionUp
			return
		}
	case directionUp:
		if canGoUp(row, column) {
			nextRow = row - 1
			nextColumn = column
			nextDirection = directionUp
			return
		} else if canGoRight(row, column) {
			nextRow = row
			nextColumn = column + 1
			nextDirection = directionRight
			return
		}
	}

	nextRow = -1
	nextColumn = -1
	return
}

func canVisited(matrix [][]int, i, j int) bool {
	if i < 0 || i >= len(matrix) {
		return false
	}

	if j < 0 || j >= len(matrix[i]) {
		return false
	}

	return matrix[i][j] != alreadyVisited
}
