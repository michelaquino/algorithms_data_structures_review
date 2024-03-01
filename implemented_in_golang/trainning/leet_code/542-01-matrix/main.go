package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/01-matrix
func main() {
	testCases := []struct {
		matrix        [][]int
		expectdOutput [][]int
	}{
		// {
		// 	matrix:        [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		// 	expectdOutput: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		// },
		// {
		// 	matrix:        [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		// 	expectdOutput: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		// },
		{
			matrix:        [][]int{{0, 0, 0, 0, 0}},
			expectdOutput: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		// {
		// 	matrix: [][]int{
		// 		{0, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		// 		{1, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		// 		{1, 1, 1, 1, 1, 0, 0, 0, 1, 1},
		// 		{1, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		// 		{0, 0, 1, 1, 1, 0, 1, 1, 1, 1},
		// 		{1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		// 		{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
		// 		{0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
		// 		{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		// 		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0},
		// 	},
		// 	expectdOutput: [][]int{
		// 		{0, 0, 1, 0, 1, 2, 1, 0, 1, 2},
		// 		{1, 1, 2, 1, 0, 1, 1, 1, 2, 3},
		// 		{2, 1, 2, 1, 1, 0, 0, 0, 1, 2},
		// 		{1, 0, 1, 0, 1, 1, 1, 0, 1, 2},
		// 		{0, 0, 1, 1, 1, 0, 1, 1, 2, 3},
		// 		{1, 0, 1, 2, 1, 1, 1, 2, 1, 2},
		// 		{1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
		// 		{0, 1, 0, 0, 0, 1, 0, 0, 1, 2},
		// 		{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		// 		{1, 0, 1, 1, 1, 0, 1, 2, 1, 0},
		// 	},
		// },
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		// nearest := bfsNearestZero(t.matrix, 1, 9)
		// fmt.Println(nearest)
		// return

		output := updateMatrix(t.matrix)
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

const (
	maxPath = (1e4) + 1
)

func updateMatrix(matrix [][]int) [][]int {
	mat := make([][]int, len(matrix))
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]int, len(matrix[0]))
	}

	queue := [][]int{}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = maxPath
				continue
			}

			queue = append(queue, []int{i, j})
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		row, column := current[0], current[1]

		neighboorsIndexes := [][]int{
			{-1, 0}, // up
			{1, 0},  // down
			{0, -1}, // left
			{0, 1},  // right
		}

		for len(neighboorsIndexes) > 0 {
			neighboorsIndex := neighboorsIndexes[0]
			neighboorsIndexes = neighboorsIndexes[1:]

			neighboorRow, neighboorColumn := row+neighboorsIndex[0], column+neighboorsIndex[1]

			if insideLimit(matrix, neighboorRow, neighboorColumn) {
				if matrix[neighboorRow][neighboorColumn] > matrix[row][column]+1 {

					matrix[neighboorRow][neighboorColumn] = matrix[row][column] + 1
					queue = append(queue, []int{neighboorRow, neighboorColumn})
				}
			}
		}

	}

	return matrix
}

func insideLimit(matrix [][]int, i, j int) bool {
	if i < 0 || j < 0 || i > len(matrix)-1 || j > len(matrix[i])-1 {
		return false
	}

	return true
}
