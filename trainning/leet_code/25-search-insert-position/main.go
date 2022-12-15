package main

import "fmt"

// https://leetcode.com/problems/search-insert-position

var (
	nums   []int
	target int
)

func main() {

	// output: 2
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))

	fmt.Println("----------------------------")
	// output: 1
	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))

	fmt.Println("----------------------------")
	// output: 4
	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))

	fmt.Println("----------------------------")
	// // output: 0
	nums = []int{1, 3, 5, 6}
	target = 0
	fmt.Println(searchInsert(nums, target))

}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for end-start >= 0 {
		middleIndex := (end + start) / 2
		middle := nums[middleIndex]

		if middle == target {
			return middleIndex
		} else if middle > target {
			end = middleIndex - 1
		} else {
			start = middleIndex + 1
		}
	}

	return end + 1
}
