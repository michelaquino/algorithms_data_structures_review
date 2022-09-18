package main

import "fmt"

// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/
func main() {
	input := []int{62, 35, 59, 55, 84, 61, 38, 87, 55, 82}

	output := findLonely(input)
	fmt.Println(output)
}

func findLonely(nums []int) []int {
	mapCount := map[int]int{}

	for _, num := range nums {
		mapCount[num] += 1
	}

	result := []int{}
	for number, count := range mapCount {
		if count > 1 {
			continue
		}

		if _, found := mapCount[number+1]; found {
			continue
		}

		if _, found := mapCount[number-1]; found {
			continue
		}

		result = append(result, number)
	}

	return result
}
