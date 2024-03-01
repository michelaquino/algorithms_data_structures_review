package main

import "fmt"

// https://leetcode.com/problems/number-of-provinces/
func main() {
	// isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	isConnected := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	// isConnected := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}
	count := findCircleNum(isConnected)
	fmt.Println("count: ", count)
}

func findCircleNum(isConnected [][]int) int {
	discovered := make([]int, len(isConnected))

	bfs := func(isConnected [][]int, start int) {
		queue := []int{start}
		for {
			if len(queue) == 0 {
				break
			}

			toProcess := queue[0]
			queue = queue[1:]

			if discovered[toProcess] == 1 {
				continue
			}

			discovered[toProcess] = 1
			for i, connected := range isConnected[toProcess] {
				if connected == 0 {
					continue
				}

				queue = append(queue, i)
			}
		}
	}

	count := 0
	for i := 0; i < len(isConnected); i++ {
		if discovered[i] != 0 {
			continue
		}

		count++
		bfs(isConnected, i)
	}

	return count
}
