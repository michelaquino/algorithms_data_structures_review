package heapsort

import "clrs_chapter_6/heap"

func Heapsort(array []int) []int {
	heap := heap.NewHeap(array, 0, len(array))
	heap.BuildMaxHeap()

	for i := heap.Length(); i >= 1; i-- {
		heap.ExchangePosition(0, i)
		heap.DownHeapSize()
		heap.MaxHeapify(0)
	}

	return heap.Array()
}
