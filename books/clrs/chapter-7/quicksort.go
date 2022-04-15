package main

import "fmt"

func main() {
	array := []int{5, 7, 2, 1, 0, 8, 3, 6, 9, 4}

	fmt.Println(array)
	arraySorted := quicksort(array, 0, len(array)-1)
	fmt.Println(arraySorted)
}

func quicksort(arrayToSort []int, startIndex, endIndex int) []int {
	array := make([]int, len(arrayToSort))
	copy(array, arrayToSort)

	if startIndex >= endIndex {
		return array
	}

	middle, arrayPartitioned := partition(array, startIndex, endIndex)
	leftSorted := quicksort(arrayPartitioned, startIndex, middle-1)
	rightSorted := quicksort(leftSorted, middle+1, endIndex)
	return rightSorted
}

func partition(array []int, startIndex, endIndex int) (int, []int) {
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)

	if startIndex < 0 {
		return startIndex, arrayCopy
	}

	if endIndex >= len(array) {
		return endIndex, arrayCopy
	}

	pivot := arrayCopy[endIndex]
	i := startIndex - 1

	for j := startIndex; j <= endIndex-1; j++ {
		if arrayCopy[j] <= pivot {
			i++
			aux := arrayCopy[i]
			arrayCopy[i] = arrayCopy[j]
			arrayCopy[j] = aux
		}
	}

	arrayCopy[endIndex] = arrayCopy[i+1]
	arrayCopy[i+1] = pivot
	return i + 1, arrayCopy
}
