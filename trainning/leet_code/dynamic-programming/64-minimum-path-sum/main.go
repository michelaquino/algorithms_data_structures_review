package main

import "fmt"

// https://leetcode.com/problems/minimum-path-sum/

func main() {
	// grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := minPathSum(grid)
	fmt.Println("result: ", result)
}

func minPathSum(grid [][]int) int {
	sumGrid := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		sumGrid[i] = make([]int, len(grid[i]))
	}

	sumGrid[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 && j > 0 {
				sumGrid[i][j] = sumGrid[i][j-1] + grid[i][j]
				continue
			}

			if j == 0 && i > 0 {
				sumGrid[i][j] = sumGrid[i-1][j] + grid[i][j]
				continue
			}

			if sumGrid[i][j-1]+grid[i][j] < sumGrid[i-1][j]+grid[i][j] {
				sumGrid[i][j] = sumGrid[i][j-1] + grid[i][j]
			} else {
				sumGrid[i][j] = sumGrid[i-1][j] + grid[i][j]
			}

		}
	}

	return sumGrid[len(sumGrid)-1][len(sumGrid[0])-1]
}

func printMatrix(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%d ", table[i][j])
		}

		fmt.Println()
	}
}
