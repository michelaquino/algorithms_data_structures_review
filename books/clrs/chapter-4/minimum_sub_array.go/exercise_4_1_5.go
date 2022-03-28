package main

import (
	"fmt"
	"math"
)

// https://algorithmist.com/wiki/Kadane%27s_algorithm
func main() {
	array := []float64{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

	maxSum := math.Inf(-1)
	maxStartIndex := -1
	maxEndIndex := -1

	var currentMax float64 = 0
	currentStartIndex := 0

	for i, value := range array {
		currentMax += value
		if currentMax > maxSum {
			maxSum = currentMax

			maxStartIndex = currentStartIndex
			maxEndIndex = i
		}

		if currentMax < 0 {
			currentMax = 0
			currentStartIndex = i + 1
		}

	}

	fmt.Printf("Max: %f; Start: %d; End: %d", maxSum, maxStartIndex, maxEndIndex)
}
