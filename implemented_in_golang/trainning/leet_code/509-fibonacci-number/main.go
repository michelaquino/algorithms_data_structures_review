package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/fibonacci-number
func main() {
	testCases := []struct {
		n             int
		expectdOutput int
	}{
		{
			n:             0,
			expectdOutput: 0,
		},
		{
			n:             1,
			expectdOutput: 1,
		},
		{
			n:             2,
			expectdOutput: 1,
		},
		{
			n:             3,
			expectdOutput: 2,
		},
		{
			n:             4,
			expectdOutput: 3,
		},
		{
			n:             10,
			expectdOutput: 55,
		},
	}

	for _, t := range testCases {

		// index := indexOfSmallerNumberThatIsGreaterThanActual([]int{2, 6, 8}, 3)
		// fmt.Println("index: ", index)

		// continue
		fmt.Println("=========================================")
		output := fib(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func fib(n int) int {
	var fibMemo func(n int, memo map[int]int) int
	fibMemo = func(n int, memo map[int]int) int {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			if result, exists := memo[n]; exists {
				return result
			}

			result := fibMemo(n-1, memo) + fibMemo(n-2, memo)
			memo[n] = result
			return result
		}
	}

	return fibMemo(n, map[int]int{})
}
