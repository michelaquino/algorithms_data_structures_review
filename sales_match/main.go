package main

import "fmt"

func main() {
	slice := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(int32(len(slice)), slice))
}

func sockMerchant(n int32, ar []int32) int32 {
	mapPairs := map[int32]int{}
	for _, a := range ar {
		mapPairs[a] += 1
	}

	var howManyPairs int32 = 0
	for _, count := range mapPairs {
		howManyPairs += int32(count / 2)
	}

	return howManyPairs
}
