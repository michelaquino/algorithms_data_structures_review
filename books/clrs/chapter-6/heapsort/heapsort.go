package heapsort

import "clrs_chapter_6/heap"

func Heapsort(array []int) []int {
	heap := heap.NewHeap(array, 0, len(array))
	buildMaxHeap(heap)

	for i := heap.Length(); i >= 1; i-- {
		heap.ExchangePosition(0, i)
		heap.DownHeapSize()
		maxHeapify(heap, 0)
	}

	return heap.Array()
}

func maxHeapify(heap *heap.Heap, position int) {
	left := left(position)
	right := right(position)
	largest := -1

	if left <= heap.Size() && heap.Array()[left] > heap.Array()[position] {
		largest = left
	} else {
		largest = position
	}

	if right <= heap.Size() && heap.Array()[right] > heap.Array()[largest] {
		largest = right
	}

	if largest != position {
		heap.ExchangePosition(largest, position)
		maxHeapify(heap, largest)
	}
}

func buildMaxHeap(heap *heap.Heap) {
	heap.SetLengthToSize()

	middle := heap.Length() / 2
	for i := middle; i >= 0; i-- {
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
