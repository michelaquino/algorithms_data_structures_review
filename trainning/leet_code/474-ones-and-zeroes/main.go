package main

import "fmt"

// https://leetcode.com/problems/ones-and-zeroes/

var (
	dp [][][]int
)

func main() {
	// strs := []string{"10", "0001", "111001", "1", "0"}
	strs := []string{"10", "0", "1"}
	zeros := 1
	ones := 1

	result := findMaxForm(strs, zeros, ones)
	fmt.Println("result: ", result)
}

func sub(quantityNumbers [][]int, index, zeros, ones int) int {
	if (zeros == 0 && ones == 0) || index == len(quantityNumbers) {
		return 0
	}

	if dp[index][zeros][ones] > -1 {
		return dp[index][zeros][ones]
	}

	if quantityNumbers[index][0] > zeros || quantityNumbers[index][1] > ones {
		ans := sub(quantityNumbers, index+1, zeros, ones)
		dp[index][zeros][ones] = ans
		return ans
	}

	include := 1 + sub(
		quantityNumbers,
		index+1,
		zeros-quantityNumbers[index][0],
		ones-quantityNumbers[index][1],
	)

	exclude := sub(
		quantityNumbers,
		index+1,
		zeros,
		ones)

	ans := max(include, exclude)
	dp[index][zeros][ones] = ans
	return ans
}

func findMaxForm(strs []string, zeros int, ones int) int {
	dp = make([][][]int, len(strs))
	for i := range dp {

		dp[i] = make([][]int, zeros+1)
		for j := range dp[i] {

			dp[i][j] = make([]int, ones+1)

			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	countZeroAndOneOnStr := func(str string) (int, int) {
		zerosOnStr := 0
		onesOnStr := 0
		for _, r := range str {
			if string(r) == "0" {
				zerosOnStr++
				continue
			}

			onesOnStr++
		}

		return zerosOnStr, onesOnStr
	}

	quantityNumbers := [][]int{}
	for _, str := range strs {
		nZero, nOne := countZeroAndOneOnStr(str)
		quantityNumbers = append(quantityNumbers, []int{nZero, nOne})
	}

	return sub(quantityNumbers, 0, zeros, ones)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
