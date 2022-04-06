package heap

import "fmt"

type Heap struct {
	array  []int
	size   int // amount of data in the array that can be considered
	length int // amount of data in the array
}

func NewHeap(array []int, startIndex, endIndex int) *Heap {
	return &Heap{
		array:  array,
		length: len(array) - 1,
		size:   endIndex - startIndex,
	}
}

func (h *Heap) ExchangePosition(position1, position2 int) {
	value1 := h.array[position1]
	h.array[position1] = h.array[position2]
	h.array[position2] = value1
}

func (h Heap) Print() {
	fmt.Println()
	for _, a := range h.array {
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}

func (h *Heap) SetLengthToSize() {
	h.size = h.length
}

func (h *Heap) DownHeapSize() {
	h.size -= 1
}

func (h Heap) Array() []int {
	return h.array
}

func (h Heap) Length() int {
	return h.length
}

func (h Heap) Size() int {
	return h.size
}
