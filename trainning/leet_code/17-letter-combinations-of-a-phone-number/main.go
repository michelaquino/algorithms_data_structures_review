package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func main() {
	testCases := []struct {
		digits        string
		expectdOutput []string
	}{
		{
			digits: "23",
			expectdOutput: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := letterCombinations(t.digits)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			fmt.Println()
			panic(fmt.Sprintf("\nexpcted %v, \ngot %v", t.expectdOutput, output))
		}
	}
}

func letterCombinations(digits string) []string {
	numberLetterMap := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
		'0': {' '},
	}

	answer := []string{}

	var backtracking func(digits string, index int, currentComb string)
	backtracking = func(digits string, index int, currentComb string) {
		if index >= len(digits) {
			if currentComb != "" {
				answer = append(answer, currentComb)
			}

			return
		}

		number := digits[index]
		for _, letter := range numberLetterMap[rune(number)] {

			currentComb += string(letter)
			backtracking(digits, index+1, currentComb)
			currentComb = currentComb[:len(currentComb)-1]
		}
	}

	backtracking(digits, 0, "")
	return answer
}
