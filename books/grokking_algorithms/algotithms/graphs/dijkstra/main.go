package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// graph := map[string]map[string]float64{
	// 	"start": {"a": 6, "b": 2},
	// 	"a":     {"end": 1},
	// 	"b":     {"a": 3, "end": 5},
	// 	"end":   {},
	// }

	// costFromStart := map[string]float64{
	// 	"a":   6,
	// 	"b":   2,
	// 	"end": math.Inf(1),
	// }

	// fathers := map[string]string{
	// 	"a":   "start",
	// 	"b":   "start",
	// 	"end": "",
	// }
	graph := map[string]map[string]float64{
		"start": {"a": 5, "b": 2},
		"a":     {"c": 4, "d": 2},
		"b":     {"a": 8, "d": 7},
		"c":     {"d": 6, "end": 3},
		"d":     {"end": 1},
		"end":   {},
	}

	costFromStart := map[string]float64{
		"a":   2,
		"b":   2,
		"c":   math.Inf(1),
		"d":   math.Inf(1),
		"end": math.Inf(1),
	}

	fathers := map[string]string{
		"a":   "start",
		"b":   "start",
		"c":   "",
		"d":   "",
		"end": "",
	}

	fathers = dijkstra(graph, costFromStart, fathers)
	printPath(fathers)
}

func printPath(fathers map[string]string) {
	path := []string{"end"}

	nextNode := fathers["end"]
	for {
		path = append([]string{nextNode}, path...)
		if nextNode == "start" {
			break
		}

		nextNode = fathers[nextNode]
	}

	fmt.Println(strings.Join(path, " -> "))
}

func dijkstra(
	graph map[string]map[string]float64,
	costFromStart map[string]float64,
	fathers map[string]string,
) map[string]string {
	alreadyProcessed := map[string]int{}

	node := findNotProcessedLowerCost(costFromStart, alreadyProcessed)
	for {
		if node == "" {
			break
		}

		cost := costFromStart[node]
		neighbors := graph[node]

		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costFromStart[n] > newCost {
				costFromStart[n] = newCost
				fathers[n] = node
			}
		}

		alreadyProcessed[node] = 1
		node = findNotProcessedLowerCost(costFromStart, alreadyProcessed)
	}

	return fathers
}

func findNotProcessedLowerCost(costs map[string]float64, processed map[string]int) string {
	lowerKey := ""
	lowerValue := math.Inf(1)
	for key, value := range costs {
		if value > lowerValue {
			continue
		}

		if _, found := processed[key]; found {
			continue
		}

		lowerValue = value
		lowerKey = key
	}

	return lowerKey
}
