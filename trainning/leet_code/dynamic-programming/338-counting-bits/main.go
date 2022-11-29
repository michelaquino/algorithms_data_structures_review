package main

import (
	"fmt"
)

// https://leetcode.com/problems/counting-bits/

func main() {
	// n := 4
	n := 15
	output := countBits(n)
	fmt.Println("output: ", output)
	fmt.Println(3 & 1)
}

// func countBits(n int) []int {
// 	ans := make([]int, n+1)
// 	ans[0] = 0
// 	if n == 0 {
// 		return ans
// 	}

// 	ans[1] = 1
// 	if n == 1 {
// 		return ans
// 	}

// 	pot := 1
// 	var twoToPot int = int(math.Pow(2, float64(pot)))
// 	calculated := 1
// 	for i := 2; i <= n; i++ {
// 		if calculated > twoToPot {
// 			calculated = 1
// 			pot++
// 			twoToPot = int(math.Pow(2, float64(pot)))
// 		}

// 		ans[i] = ans[i-twoToPot] + 1
// 		calculated++
// 	}

// 	return ans
// }

// better solution
func countBits(n int) []int {
	ones := make([]int, n+1)
	for x := 0; x <= n; x++ {
		ones[x] = ones[x>>1] + x&1
	}
	return ones
}
