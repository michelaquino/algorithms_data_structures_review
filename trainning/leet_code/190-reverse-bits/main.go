package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/house-robber
func main() {
	testCases := []struct {
		n             uint32
		expectdOutput int
	}{
		{
			n:             11,
			expectdOutput: 3,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := reverseBits(t.n)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func reverseBits(num uint32) uint32 {
	var toReturn uint32
	power := 31

	for num != 0 {
		toReturn += (num & 1) << uint32(power)
		num = num >> 1
		power -= 1
	}

	return toReturn
}
