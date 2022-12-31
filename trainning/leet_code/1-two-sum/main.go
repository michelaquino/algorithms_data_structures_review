package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/two-sum
func main() {
	testCases := []struct {
		nums          []int
		target        int
		expectdOutput []int
	}{
		{
			nums:          []int{2, 7, 11, 15},
			target:        9,
			expectdOutput: []int{0, 1},
		},
		{
			nums:          []int{3, 2, 4},
			target:        6,
			expectdOutput: []int{1, 2},
		},
		{
			nums:          []int{3, 3},
			target:        6,
			expectdOutput: []int{0, 1},
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.strs, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := twoSum(t.nums, t.target)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func twoSum(nums []int, target int) []int {
	awnser := []int{}
	hashMap := map[int]int{}

	for i, num := range nums {
		complement := target - num

		if index, exists := hashMap[complement]; exists {
			awnser = append(awnser, index, i)
		}

		hashMap[num] = i

	}

	return awnser
}
