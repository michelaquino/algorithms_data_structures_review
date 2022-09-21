package main

import "fmt"

// https://leetcode.com/problems/validate-stack-sequences/

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, len(pushed))
	top := -1
	checker := 0

	for _, i := range pushed {
		top++
		stack[top] = i

		for {
			if top < 0 { // empty
				break
			}

			if stack[top] != popped[checker] {
				break
			}

			top--
			checker++
		}
	}

	return checker == len(popped)
}
