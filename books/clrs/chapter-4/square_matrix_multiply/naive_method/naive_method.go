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

	result := naive_method(fisrt, second)
	printMatrix(result)
}

func naive_method(matrixOne, matrixTwo [][]int) [][]int {
	result := make([][]int, len(matrixOne))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(matrixOne))
	}

	for i := range matrixOne {
		for j := range matrixOne {
			for k := range matrixOne {
				result[i][j] += matrixOne[i][k] * matrixTwo[k][j]
			}
		}
	}

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
