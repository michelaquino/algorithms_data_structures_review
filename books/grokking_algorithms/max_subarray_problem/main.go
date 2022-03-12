package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}

	leftIndex, rightIndex, totalSum := findMaxSubarray(input, 0, len(input)-1)
	// leftIndex, rightIndex, totalSum := findMaxSubarrayBruteForce(input)
	fmt.Println(leftIndex)
	fmt.Println(rightIndex)
	fmt.Println(totalSum)

}

func findMaxSubarrayBruteForce(input []int) (leftIndex int, rightIndex, totalSum int) {
	diferences := buildDiferences(input)
	var max float64 = math.Inf(-1)

	for i := 0; i < len(diferences); i++ {
		var sum float64 = float64(diferences[i])
		fmt.Println()

		for j := i + 1; j < len(diferences); j++ {
			sum += float64(diferences[j])

			if sum > max {
				fmt.Printf("sum: %f; max %f\n", sum, max)
				max = sum

				fmt.Printf("i: %d => %d\n", i, diferences[i])
				fmt.Printf("j: %d => %d\n", j, diferences[j])
				leftIndex = i
				rightIndex = j
			}
		}
	}

	totalSum = int(max)
	return
}

func findMaxSubarray(input []int, lowIndex, highIndex int) (leftIndex int, rightIndex, totalSum int) {
	diferences := buildDiferences(input)

	// base case
	if lowIndex == highIndex {
		leftIndex, rightIndex = lowIndex, lowIndex
		totalSum = diferences[lowIndex]
		return
	}

	midIndex := (lowIndex + highIndex) / 2

	leftLow, leftHigh, leftSum := findMaxSubarray(input, lowIndex, midIndex)
	rightLow, rightHigh, rightSum := findMaxSubarray(input, midIndex+1, highIndex)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(input, lowIndex, midIndex, highIndex)
	if leftSum > rightHigh && leftSum > crossSum {
		leftIndex = leftLow
		rightIndex = leftHigh
		totalSum = leftSum
		return
	}

	if rightHigh > leftSum && rightHigh > crossSum {
		leftIndex = rightLow
		rightIndex = rightHigh
		totalSum = rightSum
		return
	}

	leftIndex = crossLow
	rightIndex = crossHigh
	totalSum = crossSum

	return
}

func findMaxCrossingSubarray(input []int, lowIndex, midIndex, highIndex int) (maxLeftIndex int, maxRightIndex, totalSum int) {
	diferences := buildDiferences(input)

	maxLeftIndex = -1
	maxRightIndex = -1

	var leftTotalSum float64 = math.Inf(-1)
	var sum float64 = 0
	for i := midIndex; i >= lowIndex; i-- {
		sum += float64(diferences[i])

		if sum > leftTotalSum {
			leftTotalSum = sum
			maxLeftIndex = i
		}
	}

	var rightTotalSum float64 = math.Inf(-1)
	sum = 0
	for i := midIndex + 1; i <= highIndex; i++ {
		sum += float64(diferences[i])

		if sum > rightTotalSum {
			rightTotalSum = sum
			maxRightIndex = i
		}
	}

	totalSum = int(leftTotalSum + rightTotalSum)
	return
}

func buildDiferences(input []int) []int {
	diferences := make([]int, len(input))
	for i, n := range input {
		if i == 0 {
			continue
		}

		diferences[i] = n - input[i-1]
	}

	return diferences
}
