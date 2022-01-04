package main

import (
	"fmt"
	"math"
)

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}

func miniMaxSum(arr []int32) {
	var min int64 = int64(math.Pow(10, 9))*5 + 1
	var max int64 = -1

	var sumAll int64 = 0
	for _, number := range arr {
		sumAll += int64(number)
	}

	for _, number := range arr {
		sum := sumAll - int64(number)

		if sum < min {
			min = sum
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Printf("%d %d\n", min, max)
}
