package main

import (
	"fmt"
	"reflect"
	"strings"
)

// https://leetcode.com/problems/longest-palindromic-substring
func main() {
	testCases := []struct {
		s             string
		expectdOutput string
	}{
		{
			s:             "babad",
			expectdOutput: "bab",
		},
		{
			s:             "cbbd",
			expectdOutput: "bb",
		},
	}

	for _, t := range testCases {
		// ans, l, r := lengthOfPalindrome(t.s, 2, 2)
		// fmt.Println(ans)
		// fmt.Println(l)
		// fmt.Println(r)
		// continue

		fmt.Println("=========================================")
		output := longestPalindrome(t.s)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func longestPalindrome(str string) string {
	if len(str) == 0 {
		return ""
	}

	left := -1
	right := -1
	maxLength := -1

	for i := 0; i < len(str); i++ {
		oddLength, oddLeft, oddRight := lengthOfPalindrome(str, i, i)
		evenLength, evenLeft, evenRight := lengthOfPalindrome(str, i, i+1)

		if oddLength > evenLength && oddLength > maxLength {
			maxLength = oddLength
			left = oddLeft
			right = oddRight
		} else if evenLength > oddLength && evenLength > maxLength {
			maxLength = evenLength
			left = evenLeft
			right = evenRight
		}
	}

	anwser := strings.Builder{}
	if left >= 0 && right >= 0 {
		for i := left; i <= right; i++ {
			anwser.WriteByte(str[i])
		}
	}

	return anwser.String()
}

func lengthOfPalindrome(str string, left, right int) (length, leftAws, rightAns int) {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}

	return right - left - 1, left + 1, right - 1
}
