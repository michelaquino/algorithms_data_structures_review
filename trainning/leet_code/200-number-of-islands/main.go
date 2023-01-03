package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/number-of-islands
func main() {
	testCases := []struct {
		grid          [][]byte
		expectdOutput int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expectdOutput: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expectdOutput: 3,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := numIslands(t.grid)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func printGrid(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%v ", grid[i][j])
		}

		fmt.Println()
	}
}

////////////////////////////////////////////

var (
	water      byte = '0'
	island     byte = '1'
	discovered byte = '2'
)

func numIslands(grid [][]byte) int {
	quantity := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !alreadyDiscovered(grid, i, j) && isIsland(grid, i, j) {
				quantity++
				bfs(grid, i, j)
			}
		}
	}

	return quantity
}

func bfs(grid [][]byte, i, j int) {
	if len(grid) == 0 {
		return
	}

	queue := [][]int{{i, j}}
	for len(queue) > 0 {
		toProcess := queue[0]
		queue = queue[1:]

		i, j := toProcess[0], toProcess[1]
		if !insideLimits(grid, i, j) {
			continue
		}

		if alreadyDiscovered(grid, i, j) || isWater(grid, i, j) {
			continue
		}

		grid[i][j] = discovered
		queue = append(queue,
			[]int{i + 1, j},
			[]int{i - 1, j},
			[]int{i, j + 1},
			[]int{i, j - 1},
		)
	}
}

func alreadyDiscovered(grid [][]byte, i, j int) bool {
	return grid[i][j] == discovered
}

func isWater(grid [][]byte, i, j int) bool {
	return grid[i][j] == water
}

func isIsland(grid [][]byte, i, j int) bool {
	return grid[i][j] == island
}

func insideLimits(grid [][]byte, i, j int) bool {
	if i < 0 || i > len(grid)-1 {
		return false
	}

	if j < 0 || j > len(grid[i])-1 {
		return false
	}

	return true
}
