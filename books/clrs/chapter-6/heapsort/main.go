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
		length: len(array) - 1,
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

func (h *heap) downHeapSize() {
	h.size -= 1
}

func main() {
	array := []int{2, 4, 3, 14, 1, 9, 10, 16, 8, 7}

	fmt.Println("array: ", array)
	arraySorted := heapsort(array)
	fmt.Println("arraySorted: ", arraySorted)
}

func maxHeapify(heap *heap, position int) {
	left := left(position)
	right := right(position)
	largest := -1

	if left <= heap.size && heap.array[left] > heap.array[position] {
		largest = left
	} else {
		largest = position
	}

	if right <= heap.size && heap.array[right] > heap.array[largest] {
		largest = right
	}

	if largest != position {
		heap.exchangePosition(largest, position)
		maxHeapify(heap, largest)
	}
}

func buildMaxHeap(heap *heap) {
	heap.setLengthToSize()

	middle := heap.length / 2
	for i := middle; i >= 0; i-- {
		maxHeapify(heap, i)
	}
}

func heapsort(array []int) []int {
	heap := newHeap(array, 0, len(array))
	buildMaxHeap(heap)

	for i := heap.length; i >= 1; i-- {
		heap.exchangePosition(0, i)
		heap.downHeapSize()
		maxHeapify(heap, 0)
	}

	return heap.array
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
