package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/search-a-2d-matrix
func main() {
	testCases := []struct {
		matrix        [][]int
		target        int
		expectdOutput bool
	}{
		{
			matrix:        [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:        3,
			expectdOutput: true,
		},
		{
			matrix:        [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:        13,
			expectdOutput: false,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := searchMatrix(t.matrix, t.target)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])
	return binarySearch(matrix, target, 0, (n*m)-1)
}

func binarySearch(matrix [][]int, target, start, end int) bool {
	if start > end {
		return false
	}

	n := len(matrix[0])

	pivot := (start + end) / 2
	row := pivot / n
	column := (pivot % n)

	pivotValue := matrix[row][column]
	if pivotValue == target {
		return true
	}

	if pivotValue > target {
		return binarySearch(matrix, target, start, pivot-1)
	}

	return binarySearch(matrix, target, pivot+1, end)
}
