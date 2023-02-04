package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/divisor-game/
func main() {
	testCases := []struct {
		n             int
		expectdOutput bool
	}{
		{
			n:             2,
			expectdOutput: true,
		},
		{
			n:             3,
			expectdOutput: false,
		},
		{
			n:             4,
			expectdOutput: true,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := divisorGame(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func divisorGame(n int) bool {
	return n%2 == 0
}
