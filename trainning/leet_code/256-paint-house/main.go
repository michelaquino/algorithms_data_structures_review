package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/paint-house/
func main() {
	testCases := []struct {
		costs         [][]int
		expectdOutput int
	}{
		// {
		// 	costs:         [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}},
		// 	expectdOutput: 10,
		// },
		// {
		// 	costs:         [][]int{{7, 6, 2}},
		// 	expectdOutput: 2,
		// },
		// {
		// 	costs:         [][]int{},
		// 	expectdOutput: 0,
		// },
		// {
		// 	costs:         [][]int{{5, 8, 6}, {19, 14, 13}, {7, 5, 12}, {14, 15, 17}, {3, 20, 10}},
		// 	expectdOutput: 43,
		// },
		{
			costs:         [][]int{{8, 12, 18}, {14, 6, 8}, {10, 9, 13}, {2, 17, 14}, {18, 18, 6}, {2, 1, 15}, {19, 20, 2}, {18, 15, 16}, {20, 18, 18}, {15, 10, 10}, {2, 20, 18}, {14, 5, 15}, {18, 10, 12}, {9, 17, 19}, {9, 1, 6}, {4, 4, 10}, {7, 8, 15}, {16, 17, 4}, {16, 16, 13}, {12, 7, 10}, {14, 13, 8}, {16, 6, 18}, {10, 5, 10}, {3, 5, 11}, {9, 9, 6}, {10, 15, 19}, {4, 5, 19}, {12, 17, 17}},
			expectdOutput: 43,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := minCost(t.costs)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func minCost(costs [][]int) int {
	memo := map[string]int{}

	var cost func(currentHouse, currentCostIndex int) int
	cost = func(currentHouse, currentCostIndex int) int {
		if currentHouse > len(costs)-1 {
			return 0
		}

		key := fmt.Sprintf("%d-%d", currentHouse, currentCostIndex)
		if value, exists := memo[key]; exists {
			return value
		}

		firstCost := -1
		secondCost := -1
		switch currentCostIndex {
		case 0:
			firstCost = 1
			secondCost = 2
		case 1:
			firstCost = 0
			secondCost = 2
		case 2:
			firstCost = 0
			secondCost = 1
		}

		currentCost := costs[currentHouse][currentCostIndex]
		firstPath := currentCost + cost(currentHouse+1, firstCost)
		secondPath := currentCost + cost(currentHouse+1, secondCost)
		result := min(firstPath, secondPath)
		memo[key] = result
		return result
	}

	first := cost(0, 0)
	second := cost(0, 1)
	third := cost(0, 2)
	return min(first, min(second, third))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minCost_wrong(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	colorUsed := make([]int, len(costs))
	minIndex, minValue := findMin(costs[0], -1)
	colorUsed[0] = minIndex
	minCostSum := minValue

	for i := 1; i < len(costs); i++ {
		minIndex, minValue := findMin(costs[i], colorUsed[i-1])
		colorUsed[i] = minIndex
		minCostSum += minValue
	}

	fmt.Println("colorUsed: ", colorUsed)
	return minCostSum
}

func findMin(nums []int, indexToIgnore int) (int, int) {
	minIndex := -1
	minValue := 21

	for i, num := range nums {
		if i == indexToIgnore {
			continue
		}

		if num < minValue {
			minIndex = i
			minValue = num
		}
	}

	return minIndex, minValue
}
