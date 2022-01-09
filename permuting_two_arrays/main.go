package main

import (
	"fmt"
	"math"
)

func main() {

	var k int32 = 94
	a := []int32{84, 2, 50, 51, 19, 58, 12, 90, 81, 68, 54, 73, 81, 31, 79, 85, 39, 2}
	b := []int32{53, 102, 40, 17, 33, 92, 18, 79, 66, 23, 84, 25, 38, 43, 27, 55, 8, 19}

	// var k int32 = 59
	// a := []int32{15, 16, 44, 58, 5, 10, 16, 9, 4, 20, 24}
	// b := []int32{37, 45, 41, 46, 8, 23, 59, 57, 51, 44, 59}

	fmt.Println(twoArrays(k, a, b))
}

func twoArrays(k int32, A []int32, B []int32) string {
	var greaterDifference int32 = int32(math.Pow(-10, 9)) - 1
	mapDifferenceCount := map[int32]int{}

	for _, a := range A {
		difference := k - a
		if difference > greaterDifference {
			greaterDifference = difference
		}

		mapDifferenceCount[difference] += 1
	}

	mapBCount := map[int32]int{}
	for _, b := range B {
		mapBCount[b] += 1
	}

	greaterACount := mapDifferenceCount[greaterDifference]
	bQuantityGreaterThanMaxDifference := 0
	for number, quantity := range mapBCount {
		if int32(number) >= greaterDifference {
			bQuantityGreaterThanMaxDifference += quantity
		}
	}

	if greaterACount > bQuantityGreaterThanMaxDifference {
		return "NO"
	}

	return "YES"
}

/*


-----------
5
[1, 2, 2, 1]
[3, 3, 3, 4]

[4, 3, 3, 4]

maior - 4
map {
	4: 2
	3: 2
}



*/
