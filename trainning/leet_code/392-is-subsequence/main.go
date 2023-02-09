package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/is-subsequence/
func main() {
	testCases := []struct {
		s             string
		t             string
		expectdOutput bool
	}{
		{
			s:             "abc",
			t:             "ahbgdc",
			expectdOutput: true,
		},
		{
			s:             "axc",
			t:             "ahbgdc",
			expectdOutput: false,
		},
		{
			s:             "ace",
			t:             "abcde",
			expectdOutput: true,
		},
		{
			s:             "aec",
			t:             "abcde",
			expectdOutput: false,
		},
		{
			s:             "",
			t:             "ahbgdc",
			expectdOutput: true,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := isSubsequence(t.s, t.t)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

// Faster than sequencial solution, but use more memory
func isSubsequence_recursion(s string, t string) bool {
	if s == "" {
		return true
	}

	if s != "" && t == "" {
		return false
	}

	if s[0] == t[0] {
		return isSubsequence(s[1:], t[1:])
	}

	return isSubsequence(s, t[1:])
}

// Use less memory than the recursion, but is slower
func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {

		if s[j] == t[i] {
			j++
		}
	}

	return j == len(s)
}
