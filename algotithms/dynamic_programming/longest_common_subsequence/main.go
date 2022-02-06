package main

import (
	"fmt"
	"math"
)

func main() {
	firstString := "fort"
	secondString := "fosh"

	loggestSubsequenceLength := findCommonSubsequenceLenght(firstString, secondString)
	fmt.Println("loggest common subsequence: ", loggestSubsequenceLength)
}

func findCommonSubsequenceLenght(firstString, secondString string) int {
	if firstString == "" || secondString == "" {
		return 0
	}

	matrix := make([][]int, len(firstString))
	for i := range matrix {
		matrix[i] = make([]int, len(secondString))
	}

	greatestValue := -1

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			value := 1

			// equal letters
			if firstString[i] == secondString[j] {

				if i > 0 && j > 0 {
					value = matrix[i-1][j-1] + 1
				}

				matrix[i][j] = value
				if value > greatestValue {
					greatestValue = value
				}

				continue
			}

			var upper float64 = 0
			if i > 0 {
				upper = float64(matrix[i-1][j])
			}

			var left float64 = 0
			if j > 0 {
				left = float64(matrix[i][j-1])
			}

			matrix[i][j] = int(math.Max(upper, left))
		}
	}

	// printMatrix(matrix)

	return greatestValue
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println()
	}
}
