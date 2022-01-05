package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := flippingBits(9)
	fmt.Println(result)
}

func flippingBits(n int64) int64 {
	numberInBits := strconv.FormatInt(n, 2)

	numberStr := strings.Builder{}
	numberStr.WriteString(strings.Repeat("0", 32-len(numberInBits)))
	numberStr.WriteString(numberInBits)

	reversed := strings.Builder{}
	for _, s := range numberStr.String() {
		if string(s) == "0" {
			reversed.WriteString("1")
			continue
		}

		reversed.WriteString("0")
	}

	result, _ := strconv.ParseInt(reversed.String(), 2, 64)
	return result
}
