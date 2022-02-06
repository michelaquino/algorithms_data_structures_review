package main

import "fmt"

func main() {
	var s []int32 = []int32{2, 2, 1, 3, 2}
	var d int32 = 4 // Ron's birth day - sum of segment
	var m int32 = 3 // Ron's birth month - length of segment

	fmt.Println(birthday(s, d, m))
}

func birthday(s []int32, d int32, m int32) int32 {
	var result int32 = 0
	for i := 0; i < len(s); i++ {
		subArrayEnd := i + int(m)

		if subArrayEnd > len(s) {
			continue
		}

		subarray := s[i:subArrayEnd]
		fmt.Println(subarray)
		var sum int32 = 0
		for _, number := range subarray {
			sum += number
		}

		if sum == d {
			result++
		}
	}

	return result
}
