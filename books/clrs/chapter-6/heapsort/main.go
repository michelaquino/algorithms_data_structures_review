package main

import "fmt"

type heap struct {
	array  []int
	size   int
	length int
}

func newHeap(array []int, startIndex, endIndex int) *heap {
	return &heap{
		array:  array,
		length: len(array),
		size:   endIndex - startIndex,
	}
}

func (h *heap) exchangePosition(position1, position2 int) {
	value1 := h.array[position1]
	h.array[position1] = h.array[position2]
	h.array[position2] = value1
}

func (h heap) print() {
	fmt.Println()
	for _, a := range h.array {
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}

func (h *heap) setLengthToSize() {
	h.size = h.length
}

func main() {
	array := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	heap := newHeap(array, 0, len(array))

	heap.print()
	buildMaxHeap(heap)
	heap.print()
}

func maxHeapify(heap *heap, position int) {
	left := left(position)
	right := right(position)
	largest := -1

	if left < heap.size && heap.array[left] > heap.array[position] {
		largest = left
	} else {
		largest = position
	}

	if right < heap.size && heap.array[right] > heap.array[largest] {
		largest = right
	}

	if largest != position {
		heap.exchangePosition(largest, position)
		maxHeapify(heap, largest)
	}
}

func buildMaxHeap(heap *heap) {
	heap.setLengthToSize()

	for i := len(heap.array) / 2; i >= 0; i-- {
		maxHeapify(heap, i)
	}
}

func parent(i int) int {
	return i >> 1
}

func left(i int) int {
	return (i << 1) + 1
}

func right(i int) int {
	return (i << 1) + 2
}
