package main

import "fmt"

func main() {
	// firstString := "fish"
	// secondString := "hish"
	firstString := "GeeksforGeeks”"
	secondString := "“GeeksQuiz”"

	loggestSubstringLength := findCommonSubstringLenght(firstString, secondString)
	fmt.Println("loggestSubstringLength: ", loggestSubstringLength)
}

func findCommonSubstringLenght(firstString, secondString string) int {
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

			// equal letters
			if firstString[i] == secondString[j] {

				value := 1
				if i > 0 && j > 0 {
					value = matrix[i-1][j-1] + 1
				}

				matrix[i][j] = value

				if value > greatestValue {
					greatestValue = value
				}
			}
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
