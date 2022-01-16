package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{5, 3, 1, 2, 4}
	fmt.Println(findMedian(arr))
}

func findMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	medianIndex := len(arr) / 2
	return arr[medianIndex]
}
