package main

import (
	"fmt"
)

// https://leetcode.com/problems/trapping-rain-water/
func main() {
	height := []int{5, 0, 1, 1, 1, 3}
	// height := []int{4, 2, 0, 3, 2, 5}
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 0, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	firstPointer, secondPointer := 0, 0
	result := 0

	greaterLessThanFirstPoint := 0
	positionGreater := 0
	for {
		if firstPointer > len(height)-1 {
			break
		}

		secondPointer++

		// end of array
		if secondPointer > len(height)-1 {
			// found some space that the height is less than first point
			if positionGreater > firstPointer+1 {
				for i := firstPointer + 1; i < positionGreater; i++ {
					result += greaterLessThanFirstPoint - height[i]
				}

				firstPointer, secondPointer = positionGreater, positionGreater
				greaterLessThanFirstPoint = 0
				continue
			}

			firstPointer++
			secondPointer = firstPointer
			greaterLessThanFirstPoint = 0
			continue
		}

		if height[secondPointer] > greaterLessThanFirstPoint {
			positionGreater = secondPointer
			greaterLessThanFirstPoint = height[secondPointer]
		}

		// found a heigh that can hold water
		if height[secondPointer] >= height[firstPointer] {
			lowerHeight := min(height[firstPointer], height[secondPointer])

			for i := firstPointer + 1; i < secondPointer; i++ {
				result += lowerHeight - height[i]
			}

			firstPointer = secondPointer
			greaterLessThanFirstPoint = 0
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
