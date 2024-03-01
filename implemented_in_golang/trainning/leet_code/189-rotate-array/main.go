package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/rotate-array

func main() {
	testCases := []struct {
		nums          []int
		k             int
		expectdOutput []int
	}{
		{
			nums:          []int{1, 2, 3, 4, 5, 6, 7},
			k:             3,
			expectdOutput: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:          []int{-1, -100, 3, 99},
			k:             2,
			expectdOutput: []int{3, 99, -1, -100},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		rotate(t.nums, t.k)
		fmt.Println("output: ", t.nums)

		if !reflect.DeepEqual(t.nums, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput, t.nums))
		}
	}

}

func rotate(nums []int, k int) {
	numsCopy := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newIndex := (i + k) % len(nums)
		numsCopy[newIndex] = nums[i]
	}

	copy(nums, numsCopy)

	// indexes := []int{}
	// indexes2 := []int{}
	// indexes3 := []int{}
	// for i := 0; i < len(nums); i++ {
	// 	indexes = append(indexes, (i+k)%len(nums))
	// 	fmt.Println()
	// 	fmt.Println("i: ", i)
	// 	fmt.Println("k: ", k)
	// 	fmt.Println("(i+k): ", i+k)
	// 	fmt.Println("len(nums): ", len(nums))
	// 	fmt.Println("(i+k)%len(nums): ", (i+k)%len(nums))

	// 	indexes2 = append(indexes2, (i+k+1)%len(nums))
	// 	indexes3 = append(indexes3, (i+k-1)%len(nums))
	// }

	// fmt.Println("1 - newIndex: ", indexes)
	// fmt.Println("2 - newIndex: ", indexes2)
	// fmt.Println("3 - newIndex: ", indexes3)

}
