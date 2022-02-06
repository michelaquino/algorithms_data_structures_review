package main

import (
	"fmt"
	"math"
)

func main() {
	var n int32 = 6
	var p int32 = 5
	fmt.Println(pageCount(n, p))
}

func pageCount(n int32, p int32) int32 {
	left := float64(p) / 2
	right := float64((n - p)) / 2

	if n%2 == 0 {
		return int32(math.Min(left, math.Ceil(right)))
	}

	return int32(math.Min(left, math.Floor(right)))
}
