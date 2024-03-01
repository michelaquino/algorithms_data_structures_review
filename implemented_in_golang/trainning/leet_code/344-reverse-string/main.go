package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/reverse-string

func main() {
	testCases := []struct {
		str           []byte
		expectdOutput []byte
	}{
		{
			str:           []byte("hello"),
			expectdOutput: []byte("olleh"),
		},
		{
			str:           []byte("Hannah"),
			expectdOutput: []byte("hannaH"),
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		reverseString([]byte(t.str))
		fmt.Println("output: ", []byte(t.str))

		if !reflect.DeepEqual([]byte(t.str), t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, string(t.str)))
		}
	}
}

func reverseString(str []byte) {
	start := 0
	end := len(str) - 1

	for start < end {
		aux := str[end]
		str[end] = str[start]
		str[start] = aux

		start++
		end--
	}
}
