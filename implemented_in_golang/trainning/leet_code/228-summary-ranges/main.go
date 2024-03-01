package main

import (
	"fmt"
)

// https://leetcode.com/problems/summary-ranges/
func main() {
	// nums := []int{0, 1, 2, 4, 5, 7}    // ["0->2","4->5","7"]
	nums := []int{0, 2, 3, 4, 6, 8, 9} // ["0","2->4","6","8->9"]
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	answer := []string{}

	if len(nums) == 0 {
		return answer
	}

	initialNumber := nums[0]
	finalNumber := nums[0]

	appendAwnser := func(initialNumber, finalNumber int) {
		if initialNumber == finalNumber {
			answer = append(answer, fmt.Sprintf("%d", initialNumber))
		} else {
			answer = append(answer, fmt.Sprintf("%d->%d", initialNumber, finalNumber))
		}
	}

	for i := 0; i < len(nums); i++ {
		finalNumber = nums[i]

		if i+1 < len(nums) && nums[i+1] > nums[i]+1 {
			appendAwnser(initialNumber, finalNumber)
			initialNumber = nums[i+1]
			continue
		}
	}

	appendAwnser(initialNumber, finalNumber)
	return answer
}
