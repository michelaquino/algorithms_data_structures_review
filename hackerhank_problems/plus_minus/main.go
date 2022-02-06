package main

import "fmt"

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}

func plusMinus(arr []int32) {
	var positive float32
	var negative float32
	var zero float32

	for _, number := range arr {
		if number > 0 {
			positive++
			continue
		}

		if number < 0 {
			negative++

			continue
		}

		if number == 0 {
			zero++
			continue
		}
	}

	len := float32(len(arr))
	fmt.Printf("%.6f\n", positive/len)
	fmt.Printf("%.6f\n", negative/len)
	fmt.Printf("%.6f\n", zero/len)
}
