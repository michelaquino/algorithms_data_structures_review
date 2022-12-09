package main

import "fmt"

// https://leetcode.com/problems/jump-game/
func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	// nums := []int{3, 1, 3, 0, 4}
	// nums := []int{0}
	// nums := []int{2, 0, 0}
	// nums := []int{3, 0, 8, 2, 0, 0, 1}
	// nums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0} // true

	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	fartest := 0

	for i, number := range nums {
		if fartest < i {
			return false
		}

		fartest = max(fartest, i+number)
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
