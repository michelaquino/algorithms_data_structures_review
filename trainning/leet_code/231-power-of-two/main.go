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
			n:             1,
			expectdOutput: true,
		},
		{
			n:             4,
			expectdOutput: true,
		},
		{
			n:             32,
			expectdOutput: true,
		},
		{
			n:             3,
			expectdOutput: false,
		},
		{
			n:             0,
			expectdOutput: false,
		},
		{
			n:             -1,
			expectdOutput: false,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := isPowerOfTwo(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	return n&(-n+1) == n
}

func isPowerOfTwo_2(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	return n&((^n)+1) == n
}
