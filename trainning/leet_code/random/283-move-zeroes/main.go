package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		nums          []int
		expectdOutput []int
	}{
		{
			nums:          []int{0, 1, 0, 3, 12},
			expectdOutput: []int{1, 3, 12, 0, 0},
		},
		{
			nums:          []int{0},
			expectdOutput: []int{0},
		},
		{
			nums:          []int{0, 0, 1},
			expectdOutput: []int{1, 0, 0},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		moveZeroes(t.nums)
		fmt.Println("output: ", t.nums)

		if !reflect.DeepEqual(t.nums, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput, t.nums))
		}
	}
}

func moveZeroes(nums []int) {
	i := 0
	j := 0
	for i < len(nums) {
		if nums[i] == 0 {
			i++
		} else {
			swap(nums, i, j)
			i++
			j++
		}
	}
}

func swap(nums []int, i, j int) {
	aux := nums[i]
	nums[i] = nums[j]
	nums[j] = aux
}
