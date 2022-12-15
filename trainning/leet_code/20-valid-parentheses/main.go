package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses/
func main() {

	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("["))
	fmt.Println(isValid("[["))
	fmt.Println(isValid("([[{}]])"))
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	mapParenteses := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	top := -1
	stack := []rune{}
	for _, r := range s {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
			top++
			continue
		}

		if top == -1 || stack[top] != mapParenteses[r] {
			return false
		}

		stack = stack[:top]
		top--
	}

	return top == -1
}
