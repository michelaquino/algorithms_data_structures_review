package main

import (
	"fmt"
	"math"
)

func main() {
	// m := 4 // row
	// n := 4 // column
	var r int32 = 7 //rotation factor

	// matrix := [][]int32{
	// 	{8, 9},
	// 	{14, 15},
	// 	{20, 21},
	// }
	matrix := [][]int32{
		{1, 2, 3, 4},
		{7, 8, 9, 10},
		{13, 14, 15, 16},
		{19, 20, 21, 22},
		{25, 26, 27, 28},
	}
	// matrix := [][]int32{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 16},
	// }
	// matrix := [][]int32{
	// 	{6, 7},
	// 	{10, 11},
	// }

	matrixRotation(matrix, r)
}

func matrixRotation(matrix [][]int32, r int32) {
	matrixCopy := [][]int32{}
	for _, row := range matrix {
		newRow := []int32{}
		for _, column := range row {
			newRow = append(newRow, column)
		}

		matrixCopy = append(matrixCopy, newRow)
	}

	lenRow := len(matrix)
	lenColumn := len(matrix[0])

	layers := int(math.Min(float64(lenRow), float64(lenColumn)) / 2)
	for layer := 0; layer < layers; layer++ {
		circle := []int32{}
		for i := layer; i < lenColumn-layer; i++ {
			circle = append(circle, matrixCopy[layer][i])
		}

		for i := layer + 1; i < lenRow-layer-1; i++ {
			circle = append(circle, matrixCopy[i][lenColumn-layer-1])
		}

		for i := lenColumn - layer - 1; i > layer-1; i-- {
			circle = append(circle, matrixCopy[lenRow-layer-1][i])
		}

		for i := lenRow - layer - 2; i > layer; i-- {
			circle = append(circle, matrixCopy[i][layer])
		}

		position := 0
		for i := layer; i < lenColumn-layer; i++ {
			positionRotated := (position + int(r)) % len(circle)
			matrix[layer][i] = circle[positionRotated]
			position++
		}

		for i := layer + 1; i < lenRow-layer-1; i++ {
			positionRotated := (position + int(r)) % len(circle)
			matrix[i][lenColumn-layer-1] = circle[positionRotated]
			position++
		}

		for i := lenColumn - layer - 1; i > layer-1; i-- {
			positionRotated := (position + int(r)) % len(circle)
			matrix[lenRow-layer-1][i] = circle[positionRotated]
			position++
		}

		for i := lenRow - layer - 2; i > layer; i-- {
			positionRotated := (position + int(r)) % len(circle)
			matrix[i][layer] = circle[positionRotated]
			position++
		}
	}

	printMatrix(matrix)

}

func printMatrix(matrix [][]int32) {
	for _, row := range matrix {
		for _, column := range row {
			fmt.Printf("%d ", column)
		}
		fmt.Println()
	}
}

/////////////////////////////////////////////
// This solution is not so good
func matrixRotation_1(matrix [][]int32, r int32) {
	for i := 0; i < int(r); i++ {
		rotateBorder(matrix)
	}

	printMatrix(matrix)

}

func rotateBorder(matrix [][]int32) {
	matrixCopy := [][]int32{}
	for _, row := range matrix {
		newRow := []int32{}
		for _, column := range row {
			newRow = append(newRow, column)
		}

		matrixCopy = append(matrixCopy, newRow)
	}

	lenRow := len(matrix)
	lenColumn := len(matrix[0])

	layers := int(math.Min(float64(lenRow), float64(lenColumn)) / 2)
	for layer := 0; layer < layers; layer++ {

		for i := layer; i < len(matrix)-layer; i++ {
			for j := layer; j < len(matrix[i])-layer; j++ {

				if i != layer && i != lenRow-layer-1 && j != layer && j != lenColumn-layer-1 {
					continue
				}

				newI, newJ := defineNewPosition(layer, lenRow, lenColumn, i, j)
				matrixCopyValue := matrixCopy[i][j]
				matrix[newI][newJ] = matrixCopyValue
			}
		}
	}
}

func defineNewPosition(layer, rowLengh, columnLengh, i, j int) (newI, newJ int) {
	/*
		4 situacoes

		linha 0 -> diminiu J
		linha 1 ->
			J = 0
				aumenta I
			j = 3
				diminui I
		linha 2 -> aumenta I
			J = 0
				aumenta I
			j = 3
				diminui I
		linha 3 > aumenta J

	*/

	newI = -1
	newJ = -1
	if i == layer || i == rowLengh-layer-1 {
		if i == layer {
			if j-1 < layer {

				newI = i + 1
				newJ = j
				return
			}

			newI = i
			newJ = j - 1
		} else {
			if j+1 == columnLengh-layer {
				newI = i - 1
				newJ = j
				return
			}

			newI = i
			newJ = j + 1
		}

		return
	}

	if j == layer {
		newI = i + 1
		newJ = j
		return
	}

	if j == columnLengh-layer-1 {
		newI = i - 1
		newJ = j
		return
	}

	return
}
