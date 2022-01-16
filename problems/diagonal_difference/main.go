package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	fmt.Println(diagonalDifference(matrix))

}

func diagonalDifference(arr [][]int32) int32 {
	var leftToRightDiagonal int32 = 0
	var rightToLeftDiagonal int32 = 0
	for i, row := range arr {
		leftToRightDiagonal += row[i]
		rightToLeftDiagonal += row[len(arr)-i-1]
	}

	return int32(math.Abs(float64(leftToRightDiagonal - rightToLeftDiagonal)))
}
