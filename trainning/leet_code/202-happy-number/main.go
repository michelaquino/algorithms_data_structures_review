package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/house-robber
func main() {
	testCases := []struct {
		n             int
		expectdOutput bool
	}{
		{
			n:             19,
			expectdOutput: true,
		},
		{
			n:             2,
			expectdOutput: false,
		},
		{
			n:             7,
			expectdOutput: true,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := isHappy(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	processed := map[int]bool{}
	for n != 1 {
		if _, alreadyProcessed := processed[n]; alreadyProcessed {
			return false
		}

		processed[n] = true
		numberSplited := split(n)

		sum := 0
		for _, splited := range numberSplited {
			sum += splited * splited
		}

		n = sum
	}

	return true
}

func split(n int) []int {
	result := []int{}
	for n != 0 {
		result = append(result, n%10)
		n = n / 10
	}

	return result
}
