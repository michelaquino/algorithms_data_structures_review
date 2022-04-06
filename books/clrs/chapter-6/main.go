package main

import (
	"clrs_chapter_6/heap"
	"clrs_chapter_6/heapsort"
	"fmt"
)

func main() {
	array := []int{2, 4, 3, 14, 1, 9, 10, 16, 8, 7}
	// array := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	heap := heap.NewHeap(array, 0, len(array))
	heap.BuildMaxHeap()

	max, err := heap.ExtractMax()
	if err != nil {
		panic(err)
	}

	fmt.Println("---------------------------")
	fmt.Println(max)
	heap.Print()
}

func run_heapsort() {
	array := []int{2, 4, 3, 14, 1, 9, 10, 16, 8, 7}
	fmt.Println("array: ", array)
	arraySorted := heapsort.Heapsort(array)
	fmt.Println("arraySorted: ", arraySorted)
}
