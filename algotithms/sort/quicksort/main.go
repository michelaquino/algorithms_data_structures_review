package main

import (
	"fmt"
	"time"
)

func main() {
	toSort := []int{5, 7, 2, 1, 0, 8, 3, 6, 9, 4}
	fmt.Println(toSort)

	sorted := quicksort(toSort)
	fmt.Println()
	fmt.Println(sorted)
}

func quicksort(toSort []int) []int {
	time.Sleep(1 * time.Second)
	if len(toSort) < 2 {
		return toSort
	}

	pivot := toSort[len(toSort)/2]
	lessThanPivot := []int{}
	greaterThanPivot := []int{}

	for _, number := range toSort {
		if number == pivot {
			continue
		}

		if number <= pivot {
			lessThanPivot = append(lessThanPivot, number)
			continue
		}

		greaterThanPivot = append(greaterThanPivot, number)
	}

	lessThanPivotSorted := quicksort(lessThanPivot)
	greaterThanPivotSorted := quicksort(greaterThanPivot)
	return append(append(lessThanPivotSorted, pivot), greaterThanPivotSorted...)
}
