package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/interval-list-intersections
func main() {
	testCases := []struct {
		firstList     [][]int
		secondList    [][]int
		expectdOutput [][]int
	}{
		// {
		// 	firstList:     [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		// 	secondList:    [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		// 	expectdOutput: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		// },
		// {
		// 	firstList:     [][]int{{1, 3}, {5, 9}},
		// 	secondList:    [][]int{},
		// 	expectdOutput: [][]int{},
		// },
		// {
		// 	firstList:     [][]int{{5, 10}},
		// 	secondList:    [][]int{{5, 6}},
		// 	expectdOutput: [][]int{{5, 6}},
		// },
		{
			firstList:     [][]int{{14, 16}},
			secondList:    [][]int{{7, 13}, {16, 20}},
			expectdOutput: [][]int{{16, 16}},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := intervalIntersection(t.firstList, t.secondList)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	awnser := [][]int{}
	i := 0
	j := 0

	for i <= len(firstList)-1 && j <= len(secondList)-1 {
		lowest := max(firstList[i][0], secondList[j][0])
		highest := min(firstList[i][1], secondList[j][1])

		if lowest <= highest {
			awnser = append(awnser, []int{lowest, highest})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return awnser
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
