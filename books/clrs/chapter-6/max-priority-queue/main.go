package main

import (
	"errors"
	"fmt"
)

type heap struct {
	array      []int
	endIndex   int
	startIndex int
}

func newHeap(array []int, startIndex, endIndex int) *heap {
	return &heap{
		array:      array,
		startIndex: startIndex,
		endIndex:   endIndex,
	}
}

// amount of data in the array that can be considered
func (h heap) size() int {
	return h.endIndex - h.startIndex
}

// amount of data in the array
func (h heap) length() int {
	return len(h.array) - 1
}

func (h heap) print() {
	fmt.Println()
	for _, a := range h.array {
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}

func (h heap) maximum() int {
	return h.array[0]
}

func (h heap) extractMax() (int, error) {
	if h.size() < 1 {
		return -1, errors.New("heap underflow")
	}

	return 0, nil
}
