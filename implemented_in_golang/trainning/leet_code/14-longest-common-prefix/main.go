package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/longest-common-prefix
func main() {
	testCases := []struct {
		strs          []string
		expectdOutput string
	}{
		{
			strs:          []string{"flower", "flow", "flight"},
			expectdOutput: "fl",
		},
		{
			strs:          []string{"dog", "racecar", "car"},
			expectdOutput: "",
		},
		{
			strs:          []string{"flower", "flower", "flower", "flower"},
			expectdOutput: "flower",
		},
		{
			strs:          []string{"cir", "car"},
			expectdOutput: "c",
		},
		{
			strs:          []string{"bab", "bcc"},
			expectdOutput: "b",
		},
		{
			strs:          []string{"aca", "cba"},
			expectdOutput: "",
		},
		{
			strs:          []string{"a"},
			expectdOutput: "a",
		},
		{
			strs:          []string{"babb", "caa"},
			expectdOutput: "",
		},
		{
			strs:          []string{"a", "b"},
			expectdOutput: "",
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.strs, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := longestCommonPrefix(t.strs)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func longestCommonPrefix(strs []string) string {
	lessStrLen := 201
	lessStrLenIndex := -1
	for i, s := range strs {
		if len(s) < lessStrLen {
			lessStrLen = len(s)
			lessStrLenIndex = i
		}
	}

	if lessStrLen == 0 || lessStrLenIndex == -1 {
		return ""
	}

	firstPointer := 0
	secondPointer := 0
	foundADifferentLetter := false
	for secondPointer < lessStrLen {
		letter := strs[lessStrLenIndex][secondPointer]

		for _, s := range strs {
			if s[secondPointer] != letter {
				foundADifferentLetter = true
				break
			}
		}

		if foundADifferentLetter {
			break
		}

		secondPointer++
	}

	secondPointer--
	maxStr := ""
	for i := firstPointer; i <= secondPointer; i++ {
		maxStr += string(strs[0][i])
	}

	return maxStr
}
