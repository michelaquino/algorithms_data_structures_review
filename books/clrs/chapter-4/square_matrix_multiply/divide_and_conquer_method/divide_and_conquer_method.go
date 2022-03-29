package main

import "fmt"

func main() {
	fisrt := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	second := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	// fisrt := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// second := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	result := divideAndConquerMethod(fisrt, second)
	fmt.Println()
	printMatrix(result)
}

func divideAndConquerMethod(first, second [][]int) [][]int {
	result := newEmptyMatrix(len(first))
	if len(first) == 1 {
		result[0][0] = first[0][0] * second[0][0]
		return result
	}

	size := len(first)
	newSize := int(size / 2)

	firstTopLeft := createMatrixFrom(first, 0, newSize-1, 0, newSize-1)
	firstTopRight := createMatrixFrom(first, 0, newSize-1, newSize, size-1)
	firstBottonLeft := createMatrixFrom(first, newSize, size-1, 0, newSize-1)
	firstBottonRight := createMatrixFrom(first, newSize, size-1, newSize, size-1)

	secondTopLeft := createMatrixFrom(second, 0, newSize-1, 0, newSize-1)
	secondTopRight := createMatrixFrom(second, 0, newSize-1, newSize, size-1)
	secondBottonLeft := createMatrixFrom(second, newSize, size-1, 0, newSize-1)
	secondBottonRight := createMatrixFrom(second, newSize, size-1, newSize, size-1)

	topLeftResult := sum(divideAndConquerMethod(firstTopLeft, secondTopLeft), divideAndConquerMethod(firstTopRight, secondBottonLeft))
	topRightResult := sum(divideAndConquerMethod(firstTopLeft, secondTopRight), divideAndConquerMethod(firstTopRight, secondBottonRight))
	bottonLeftResult := sum(divideAndConquerMethod(firstBottonLeft, secondTopLeft), divideAndConquerMethod(firstBottonRight, secondBottonLeft))
	bottonRightResult := sum(divideAndConquerMethod(firstBottonLeft, secondTopRight), divideAndConquerMethod(firstBottonRight, secondBottonRight))

	result = copyTo(topLeftResult, result, 0, newSize-1, 0, newSize-1)
	result = copyTo(topRightResult, result, 0, newSize-1, newSize, size-1)
	result = copyTo(bottonLeftResult, result, newSize, size-1, 0, newSize-1)
	result = copyTo(bottonRightResult, result, newSize, size-1, newSize, size-1)

	return result
}

func printMatrix(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%d ", table[i][j])
		}

		fmt.Println()
	}
}

func sum(first, second [][]int) [][]int {
	result := newEmptyMatrix(len(first))

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(first); j++ {
			result[i][j] = first[i][j] + second[i][j]
		}
	}

	return result
}

func newEmptyMatrix(size int) [][]int {
	result := make([][]int, size)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, size)
	}

	return result
}

func createMatrixFrom(origin [][]int, startRowIndex, endRowIndex, startColumnIndex, endColumnIndex int) [][]int {
	size := endRowIndex - startRowIndex + 1

	copied := make([][]int, size)
	for i := 0; i < size; i++ {
		copied[i] = make([]int, size)
	}

	i := 0
	for r := startRowIndex; r <= endRowIndex; r++ {
		j := 0
		for c := startColumnIndex; c <= endColumnIndex; c++ {
			copied[i][j] = origin[r][c]
			j++
		}

		i++
	}

	return copied
}

func copyTo(from [][]int, to [][]int, startRowIndex, endRowIndex, startColumnIndex, endColumnIndex int) [][]int {
	i := 0
	for r := startRowIndex; r <= endRowIndex; r++ {
		j := 0
		for c := startColumnIndex; c <= endColumnIndex; c++ {
			to[r][c] = from[i][j]
			j++
		}

		i++
	}

	return to
}
