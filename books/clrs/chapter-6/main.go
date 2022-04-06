package main

import (
	"clrs_chapter_6/heapsort"
	"fmt"
)

func main() {
	array := []int{2, 4, 3, 14, 1, 9, 10, 16, 8, 7}

	fmt.Println("array: ", array)
	arraySorted := heapsort.Heapsort(array)
	fmt.Println("arraySorted: ", arraySorted)
}
