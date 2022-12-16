package main

import (
	"fmt"
	"reflect"
)

var (
	notProcessed              = 0
	pieceOfIslandNotProcessed = 1
	alreadyProcessed          = -1
	islandAlreadyFound        = 2
)

// https://leetcode.com/problems/max-area-of-island
func main() {
	testCases := []struct {
		grid          [][]int
		expectdOutput int
	}{
		{
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			expectdOutput: 6,
		},
		{
			grid:          [][]int{{0, 0, 0, 0, 0, 0, 0, 0}},
			expectdOutput: 0,
		},
		{
			grid: [][]int{ // biggest island starts on (2, 33)
				{0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0},
				{0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1},
				{0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1},
				{1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1},
			},
			expectdOutput: 14,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := maxAreaOfIsland(t.grid)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func maxAreaOfIsland(grid [][]int) int {
	result := dfs(grid, 0, 0)
	return result
}

func dfs(grid [][]int, i, j int) int {
	if !insideLimits(grid, i, j) {
		return 0
	}

	if grid[i][j] == alreadyProcessed {
		return 0
	}

	maxLength := 0
	if isPieceOfIsland(grid, i, j) {
		maxLength = max(maxLength, legthOfIslandBFS(grid, i, j))
	}

	grid[i][j] = alreadyProcessed
	maxLength = max(maxLength, dfs(grid, i+1, j))
	maxLength = max(maxLength, dfs(grid, i-1, j))
	maxLength = max(maxLength, dfs(grid, i, j+1))
	maxLength = max(maxLength, dfs(grid, i, j-1))

	return maxLength
}

func legthOfIslandBFS(grid [][]int, startI, startJ int) int {
	queue := [][]int{{startI, startJ}}

	length := 0
	for len(queue) > 0 {
		toProcess := queue[0]
		queue = queue[1:]
		i, j := toProcess[0], toProcess[1]

		if !isPieceOfIsland(grid, i, j) {
			continue
		}

		if isPieceOfIslandAlreadyProcessed(grid, i, j) {
			continue
		}

		grid[i][j] = islandAlreadyFound

		// up
		if validNeighboor(grid, i-1, j) {
			queue = append(queue, []int{i - 1, j})
		}

		// down
		if validNeighboor(grid, i+1, j) {
			queue = append(queue, []int{i + 1, j})
		}

		// left
		if validNeighboor(grid, i, j-1) {
			queue = append(queue, []int{i, j - 1})
		}

		// right
		if validNeighboor(grid, i, j+1) {
			queue = append(queue, []int{i, j + 1})
		}

		length++
	}

	return length
}

func validNeighboor(grid [][]int, i, j int) bool {
	return insideLimits(grid, i, j) && isPieceOfIsland(grid, i, j)
}

func isPieceOfIsland(grid [][]int, i, j int) bool {
	return grid[i][j] == pieceOfIslandNotProcessed || grid[i][j] == islandAlreadyFound
}

func isPieceOfIslandAlreadyProcessed(grid [][]int, i, j int) bool {
	return grid[i][j] == islandAlreadyFound
}

func insideLimits(grid [][]int, i, j int) bool {
	if i < 0 || i > len(grid)-1 {
		return false
	}

	if j < 0 || j > len(grid[i])-1 {
		return false
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func printMatrix(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}

		fmt.Println()
	}
}
