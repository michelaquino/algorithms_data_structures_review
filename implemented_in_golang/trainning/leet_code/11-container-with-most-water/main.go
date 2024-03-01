package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/container-with-most-water
func main() {
	testCases := []struct {
		height        []int
		expectdOutput int
	}{
		{
			height:        []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expectdOutput: 49,
		},
		{
			height:        []int{1, 1},
			expectdOutput: 1,
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.strs, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := maxArea(t.height)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxVolumn := -1
	for left < right {
		width := right - left
		minHeigth := min(height[left], height[right])
		maxVolumn = max(maxVolumn, minHeigth*width)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolumn
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
