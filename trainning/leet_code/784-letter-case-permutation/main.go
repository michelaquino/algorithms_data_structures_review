package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/letter-case-permutation
func main() {
	testCases := []struct {
		s             string
		expectdOutput []string
	}{
		// {
		// 	s:             "abc",
		// 	expectdOutput: []string{},
		// },
		{
			s:             "a1b2",
			expectdOutput: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
		},
		{
			s:             "3z4",
			expectdOutput: []string{"3z4", "3Z4"},
		},
		{
			s:             "C",
			expectdOutput: []string{"C", "c"},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := letterCasePermutation(t.s)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func letterCasePermutation(s string) []string {
	var awnser = []string{}

	var backtracking func(originStr string, actualPosition int, result string)
	backtracking = func(originStr string, actualPosition int, result string) {
		if actualPosition >= len(originStr) {
			awnser = append(awnser, result)
			return
		}

		upperCaseOptions := []bool{false, true}
		char := rune(originStr[actualPosition])

		for _, isUpperCase := range upperCaseOptions {
			toIncrement := string(char)
			if isUpperCase {
				if !unicode.IsLetter(char) {
					continue
				}

				if unicode.IsUpper(char) {
					toIncrement = strings.ToLower(string(char))
				} else {
					toIncrement = strings.ToUpper(string(char))
				}

			}

			result = result + toIncrement
			backtracking(originStr, actualPosition+1, result)
			result = result[:len(result)-1]
		}
	}

	backtracking(s, 0, "")
	return awnser
}
