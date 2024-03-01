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

// https://leetcode.com/problems/unique-paths
func main() {
	testCases := []struct {
		m             int
		n             int
		expectdOutput int
	}{
		{
			m:             3,
			n:             2,
			expectdOutput: 3,
		},
		{
			m:             3,
			n:             7,
			expectdOutput: 28,
		},
	}

	for _, t := range testCases {
		// ans := findPath(t.graph, 0, 3)
		// fmt.Println(ans)
		// continue

		fmt.Println("=========================================")
		output := uniquePaths(t.m, t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func uniquePaths(m int, n int) int {
	return dunamicProgramming(m, n, 0, 0, map[string]int{})
}

func dunamicProgramming(m, n, actualN, actualM int, memo map[string]int) int {
	key := fmt.Sprintf("%d-%d", actualN, actualM)
	if result, exists := memo[key]; exists {
		return result
	}

	if actualM > m-1 || actualN > n-1 {
		memo[key] = 0
		return 0
	}

	if actualM == m-1 && actualN == n-1 {
		memo[key] = 1
		return 1
	}

	left := dunamicProgramming(m, n, actualN, actualM+1, memo)
	right := dunamicProgramming(m, n, actualN+1, actualM, memo)

	result := left + right
	memo[key] = result
	return result
}

func printMatrix(graph [][]int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Printf("%d ", graph[i][j])
		}

		fmt.Println()
	}
}
