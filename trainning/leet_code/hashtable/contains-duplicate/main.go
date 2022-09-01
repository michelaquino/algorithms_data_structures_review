package main

import "fmt"

// https://leetcode.com/problems/contains-duplicate/

func main() {
	testOne := []int{1, 2, 3, 1}
	testTwo := []int{1, 2, 3, 4}
	testThree := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println("testOne: ", containsDuplicate(testOne))
	fmt.Println("testTwo: ", containsDuplicate(testTwo))
	fmt.Println("testThree: ", containsDuplicate(testThree))

}

func containsDuplicate(nums []int) bool {
	numbersMap := map[int]int{}

	for _, num := range nums {
		numbersMap[num]++
		count := numbersMap[num]

		if count >= 2 {
			return true
		}
	}

	return false
}
