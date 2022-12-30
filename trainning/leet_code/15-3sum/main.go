package main

import (
	"fmt"
	"reflect"
	"sort"
)

// https://leetcode.com/problems/3sum
func main() {
	testCases := []struct {
		nums          []int
		expectdOutput [][]int
	}{
		{
			nums:          []int{-1, 0, 1, 2, -1, -4},
			expectdOutput: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:          []int{0, 1, 1},
			expectdOutput: [][]int{},
		},
		{
			nums:          []int{0, 0, 0},
			expectdOutput: [][]int{{0, 0, 0}},
		},
	}

	for _, t := range testCases {
		// result := findBallOnColumn(t.strs, 0)
		// fmt.Println("result: ", result)
		// return

		fmt.Println("=========================================")
		output := threeSum(t.nums)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func threeSum(nums []int) [][]int {
	answer := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return answer
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lowerPointer := i + 1
		higherPointer := len(nums) - 1

		for lowerPointer < higherPointer {
			sum := nums[i] + nums[lowerPointer] + nums[higherPointer]
			if sum < 0 {
				lowerPointer++
			} else if sum > 0 {
				higherPointer--
			} else {
				answer = append(answer, []int{nums[i], nums[lowerPointer], nums[higherPointer]})
				lowerPointer++
				higherPointer--

				for lowerPointer < higherPointer && nums[lowerPointer] == nums[lowerPointer-1] {
					lowerPointer++
				}
			}
		}
	}

	return answer
}
