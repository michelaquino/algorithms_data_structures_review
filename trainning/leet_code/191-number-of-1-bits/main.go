package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/house-robber
func main() {
	testCases := []struct {
		n             uint32
		expectdOutput int
	}{
		{
			n:             11,
			expectdOutput: 3,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := hammingWeight(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}

		num = num >> 1
	}

	return count
}
