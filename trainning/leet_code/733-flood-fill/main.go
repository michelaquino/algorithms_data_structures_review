package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/flood-fill
func main() {
	testCases := []struct {
		image         [][]int
		sr            int
		sc            int
		color         int
		expectdOutput [][]int
	}{
		{
			image:         [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:            1,
			sc:            1,
			color:         2,
			expectdOutput: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image:         [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:            0,
			sc:            0,
			color:         0,
			expectdOutput: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := floodFill(t.image, t.sr, t.sc, t.color)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

// BFS
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// printImage(image)

	originalColor := image[sr][sc]
	queue := [][]int{{sr, sc}}

	// key: i;j
	processedMap := map[string]bool{}
	for len(queue) > 0 {
		toProcess := queue[0]
		queue = queue[1:]
		i, j := toProcess[0], toProcess[1]
		mapKey := fmt.Sprintf("%d;%d", i, j)

		if alreadyProcessed := processedMap[mapKey]; alreadyProcessed {
			continue
		}

		if image[i][j] != originalColor {
			continue
		}

		image[i][j] = color
		processedMap[mapKey] = true

		// up
		if insideImageLimit(image, i-1, j) {
			queue = append(queue, []int{i - 1, j})
		}

		// down
		if insideImageLimit(image, i+1, j) {
			queue = append(queue, []int{i + 1, j})
		}

		// left
		if insideImageLimit(image, i, j-1) {
			queue = append(queue, []int{i, j - 1})
		}

		// right
		if insideImageLimit(image, i, j+1) {
			queue = append(queue, []int{i, j + 1})
		}
	}

	return image
}

func insideImageLimit(image [][]int, i, j int) bool {
	if i < 0 || i > len(image)-1 {
		return false
	}

	if j < 0 || j > len(image[i])-1 {
		return false
	}

	return true
}

func printImage(image [][]int) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			fmt.Printf("%d ", image[i][j])
		}

		fmt.Println()
	}
}
