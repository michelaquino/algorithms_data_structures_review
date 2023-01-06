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

// https://leetcode.com/problems/all-paths-from-source-to-target
func main() {
	testCases := []struct {
		graph         [][]int
		expectdOutput [][]int
	}{
		// {
		// 	graph: [][]int{
		// 		{1, 2},
		// 		{3},
		// 		{3},
		// 		{},
		// 	},
		// 	expectdOutput: [][]int{
		// 		{0, 1, 3},
		// 		{0, 2, 3},
		// 	},
		// },
		// {
		// 	graph: [][]int{ // biggest island starts on (2, 33)
		// 		{4, 3, 1},
		// 		{3, 2, 4},
		// 		{3},
		// 		{4},
		// 		{},
		// 	},
		// 	expectdOutput: [][]int{
		// 		{0, 4},
		// 		{0, 3, 4},
		// 		{0, 1, 3, 4},
		// 		{0, 1, 2, 3, 4},
		// 		{0, 1, 4},
		// 	},
		// },
		{
			graph: [][]int{ // biggest island starts on (2, 33)
				{3, 1},
				{4, 6, 7, 2, 5},
				{4, 6, 3},
				{6, 4},
				{7, 6, 5},
				{6},
				{7},
				{},
			},
			expectdOutput: [][]int{
				{0, 3, 6, 7},
				{0, 3, 4, 7},
				{0, 3, 4, 6, 7},
				{0, 3, 4, 5, 6, 7},
				{0, 1, 4, 7},
				{0, 1, 4, 6, 7},
				{0, 1, 4, 5, 6, 7},
				{0, 1, 6, 7},
				{0, 1, 7},
				{0, 1, 2, 4, 7},
				{0, 1, 2, 4, 6, 7},
				{0, 1, 2, 4, 5, 6, 7},
				{0, 1, 2, 6, 7},
				{0, 1, 2, 3, 6, 7},
				{0, 1, 2, 3, 4, 7},
				{0, 1, 2, 3, 4, 6, 7},
				{0, 1, 2, 3, 4, 5, 6, 7},
				{0, 1, 5, 6, 7},
			},
		},
	}

	for _, t := range testCases {
		// ans := findPath(t.graph, 0, 3)
		// fmt.Println(ans)
		// continue

		fmt.Println("=========================================")
		output := allPathsSourceTarget(t.graph)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	var findPath func(graph [][]int, node, target int, currentPath []int)

	findPath = func(graph [][]int, node, target int, currentPath []int) {
		if node > len(graph) {
			return
		}

		currentPathCopy := make([]int, len(currentPath))
		copy(currentPathCopy, currentPath)

		currentPathCopy = append(currentPathCopy, node)
		if node == target {
			result = append(result, currentPathCopy)
			return
		}

		for _, edge := range graph[node] {
			findPath(graph, edge, target, currentPathCopy)
		}
	}

	target := len(graph) - 1
	findPath(graph, 0, target, []int{})
	return result
}

// func findPath(graph [][]int, node, target int) [][]int {
// 	fmt.Println("================")
// 	fmt.Println("=> node: ", node)
// 	fmt.Println("=> target: ", target)
// 	if node > len(graph) {
// 		// fmt.Println("==> 1")
// 		return [][]int{}
// 	}

// 	ans := [][]int{}
// 	for _, edge := range graph[node] {
// 		fmt.Println("---")
// 		fmt.Println("=> ans: ", ans)
// 		fmt.Println("=> node: ", node)
// 		fmt.Println("edge: ", edge)

// 		if edge == target {
// 			ans = append([][]int{{node, edge}})
// 			continue
// 		}

// 		nextPath := findPath(graph, edge, target)
// 		fmt.Println("nextPath: ", nextPath)
// 		if len(nextPath) != 0 {

// 			r := [][]int{}
// 			for _, n := range nextPath {
// 				completeNext := append([]int{node}, n...)
// 				r = append(r, completeNext)
// 			}

// 			// nextPath = append([]int{node}, nextPath...)
// 			ans = append(ans, r...)
// 		}
// 	}

// 	return ans
// }

func printMatrix(graph [][]int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Printf("%d ", graph[i][j])
		}

		fmt.Println()
	}
}
