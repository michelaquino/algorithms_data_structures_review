package main

import "fmt"

func main() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(lonelyinteger(a))
}

func lonelyinteger(a []int32) int32 {
	mapCount := map[int32]int32{}

	for _, number := range a {
		mapCount[int32(number)] = mapCount[int32(number)] + 1
	}

	for number, count := range mapCount {
		if count == 1 {
			return number
		}
	}

	return 0
}
