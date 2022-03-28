package main

import "fmt"

func main() {
	matrixOne := [][]int{
		{1, 2},
		{3, 4},
	}

	matrixTwo := [][]int{
		{5, 6},
		{7, 8},
	}

	result := naive_method(matrixOne, matrixTwo)
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
				result[i][j] = matrixOne[i][k] * matrixTwo[k][j]
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
