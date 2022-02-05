package main

import (
	"fmt"
)

func main() {
	toSort := []int{5, 7, 2, 1, 0, 8, 3, 6, 9, 4}
	sorted := mergesort(toSort)
	fmt.Println()
	fmt.Println(sorted)
}

func mergesort(toSort []int) []int {
	if len(toSort) < 2 {
		return toSort
	}

	middle := len(toSort) / 2
	left := toSort[:middle]
	right := toSort[middle:]

	leftSorted := mergesort(left)
	rightSorted := mergesort(right)

	return merge(leftSorted, rightSorted)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}

		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
