package main

import (
	"fmt"
)

//https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
// https://math.stackexchange.com/questions/666288/number-of-ways-to-interleave-two-ordered-sequences/666295#666295

func printTriangle(triangle [][]int) {
	for i := 0; i < len(triangle); i++ {
		fmt.Println("")

		for j := 0; j < len(triangle[i]); j++ {
			fmt.Printf("%d\t", triangle[i][j])
		}
	}
}

var (
	modulo = int(1e9) + 7
)

func main() {
	nums := []int{19, 3, 57, 34, 15, 89, 58, 35, 2, 33, 46, 13, 40, 79, 60, 30, 61, 26, 54, 22, 84, 51, 75, 6, 87, 44, 55, 48, 27, 8, 72, 47, 16, 69, 36, 76, 41, 1, 80, 62, 73, 24, 93, 50, 92, 65, 39, 5, 32, 67, 12, 29, 90, 45, 9, 38, 88, 52, 10, 85, 74, 66, 83, 18, 20, 77, 49, 28, 23, 53, 86, 64, 78, 82, 37, 42, 56, 17, 81, 4, 14, 70, 59, 31, 7, 25, 43, 68, 91, 71, 21, 63, 94, 11}

	// nums := []int{31, 23, 14, 24, 15, 12, 25, 28, 5, 35, 17, 6, 9, 11, 1, 27, 18, 20, 2, 3, 33, 10, 13, 4, 7, 36, 32, 29, 8, 30, 26, 19, 34, 22, 21, 16}
	// nums := []int{9, 4, 2, 1, 3, 6, 5, 7, 8, 14, 11, 10, 12, 13, 16, 15, 17, 18}
	// nums := []int{2, 1, 3}
	// nums := []int{6, 5, 7, 8}
	ways := numOfWays(nums)
	fmt.Println(ways)
}

func numOfWays(nums []int) int {
	triangle := pascalTriangle(len(nums) + 1)
	ways := calcNumOfWays(triangle, nums) - 1
	return ways % modulo
}

func calcNumOfWays(pascalTriangle [][]int, nums []int) int {
	if len(nums) <= 2 {
		return 1
	}

	left, right := getLessAndGreaterThan(nums)
	waysLeft := calcNumOfWays(pascalTriangle, left) % modulo
	waysRight := calcNumOfWays(pascalTriangle, right) % modulo

	permutations := pascalTriangle[len(left)+len(right)][len(right)] % modulo
	totalPermutations := (permutations * waysLeft) % modulo
	totalPermutations = (totalPermutations * waysRight) % modulo

	return totalPermutations
}

func getLessAndGreaterThan(numbers []int) ([]int, []int) {
	left := []int{}
	right := []int{}

	number := numbers[0]
	for _, n := range numbers {
		if n < number {
			left = append(left, n)
			continue
		}

		if n > number {
			right = append(right, n)
		}
	}

	return left, right
}

func pascalTriangle(lenght int) [][]int {
	triangle := make([][]int, lenght)

	for i := 0; i < lenght; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = (triangle[i-1][j-1] + triangle[i-1][j]) % modulo
		}
	}

	return triangle
}
