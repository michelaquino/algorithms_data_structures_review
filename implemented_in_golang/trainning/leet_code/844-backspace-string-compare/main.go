package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/backspace-string-compare
func main() {
	testCases := []struct {
		s             string
		t             string
		expectdOutput bool
	}{
		{
			s:             "ab#c",
			t:             "ad#c",
			expectdOutput: true,
		},
		{
			s:             "ab##",
			t:             "c#d#",
			expectdOutput: true,
		},
		{
			s:             "a#c",
			t:             "b",
			expectdOutput: false,
		},
		{
			s:             "a##c",
			t:             "#a#c",
			expectdOutput: true,
		},
		{
			s:             "bxj##tw",
			t:             "bxo#j##tw",
			expectdOutput: true,
		},
		{
			s:             "y#fo##f",
			t:             "y#fx#o##f",
			expectdOutput: true,
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.strs, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := backspaceCompare(t.s, t.t)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func backspaceCompare(s string, t string) bool {
	pointerS := len(s) - 1
	skipS := 0

	pointerT := len(t) - 1
	skipT := 0

	for pointerS >= 0 || pointerT >= 0 {

		for pointerS >= 0 {
			if string(s[pointerS]) == "#" {
				skipS++
				pointerS--
			} else if skipS > 0 {
				skipS--
				pointerS--
			} else {
				break
			}
		}

		for pointerT >= 0 {
			if string(t[pointerT]) == "#" {
				skipT++
				pointerT--
			} else if skipT > 0 {
				skipT--
				pointerT--
			} else {
				break
			}
		}

		if pointerS >= 0 && pointerT >= 0 && s[pointerS] != t[pointerT] {
			return false
		}

		if (pointerS >= 0) != (pointerT >= 0) {
			return false
		}

		pointerS--
		pointerT--
	}

	return true
}

func backspaceCompare_wrong(s string, t string) bool {
	first := len(s) - 1
	second := len(t) - 1

	resultS := ""
	resultT := ""

	for first >= 0 || second >= 0 {
		// fmt.Println()
		// fmt.Println("resultS: ", resultS)
		// fmt.Println("resultT: ", resultT)
		// fmt.Println("first: ", first)
		// fmt.Println("second: ", second)

		if first >= 0 {
			if string(s[first]) == "#" {
				aux := first

				for aux >= 0 && string(s[aux]) == "#" {
					first -= 2
					aux--
				}
			} else {
				if first < len(s)-1 && string(s[first+1]) == "#" {
					first -= 2
				}

				if first >= 0 {
					resultS = string(s[first]) + resultS
				}

				first--
			}
		}

		if second >= 0 {
			if string(t[second]) == "#" {
				aux := second

				for aux >= 0 && string(t[aux]) == "#" {
					second -= 2
					aux--
				}
			} else {
				if second < len(t)-1 && string(t[second+1]) == "#" {
					second -= 2
				}

				if second >= 0 {
					resultT = string(t[second]) + resultT
				}

				second--
			}
		}
	}

	// fmt.Println()
	// fmt.Println("resultS: ", resultS)
	// fmt.Println("resultT: ", resultT)
	// fmt.Println("first: ", first)
	// fmt.Println("second: ", second)

	return resultS == resultT
}
