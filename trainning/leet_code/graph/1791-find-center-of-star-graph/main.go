package main

import "fmt"

// https://leetcode.com/problems/find-center-of-star-graph/

func main() {
	// edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	edges := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}

	fmt.Println("center: ", findCenter(edges))
}

func findCenter(edges [][]int) int {
	countEdges := map[int]int{}
	result := -1

	maxCount := -1
	for _, edge := range edges {
		for _, vertex := range edge {
			countEdges[vertex] += 1
			count := countEdges[vertex]

			if count > maxCount {
				maxCount = count
				result = vertex
			}
		}
	}

	return result
}
