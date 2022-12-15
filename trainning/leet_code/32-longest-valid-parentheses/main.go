package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-valid-parentheses/

func main() {
	// fmt.Println(longestValidParentheses("(()"))
	// fmt.Println(longestValidParentheses(")()())"))
	// fmt.Println(longestValidParentheses("()(()"))
	// fmt.Println(longestValidParentheses("())((())"))
	fmt.Println(longestValidParentheses(")("))
}

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	maxLenght := 0
	leftBrackets := 0
	rightBrackets := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftBrackets++
		} else {
			rightBrackets++
		}

		if leftBrackets == rightBrackets && leftBrackets*2 > maxLenght {
			maxLenght = leftBrackets * 2
		} else if rightBrackets > leftBrackets {
			leftBrackets, rightBrackets = 0, 0
		}
	}

	leftBrackets, rightBrackets = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			leftBrackets++
		} else {
			rightBrackets++
		}

		if leftBrackets == rightBrackets && leftBrackets*2 > maxLenght {
			maxLenght = leftBrackets * 2
		} else if leftBrackets > rightBrackets {
			leftBrackets, rightBrackets = 0, 0
		}
	}

	return maxLenght
}

func printStack(stack []rune) {
	fmt.Printf("--> ")
	for _, r := range stack {
		fmt.Printf("%s ", string(r))
	}
	fmt.Println()
}
