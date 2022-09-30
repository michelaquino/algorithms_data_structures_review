package main

import (
	"fmt"
)

/*
Insert
Find min
Delete Min
*/

func main() {
	h := newHeap()
	h.insert(50)
	fmt.Println("h.data: ", h.data)
	h.insert(30)
	fmt.Println("h.data: ", h.data)
	h.insert(60)
	fmt.Println("h.data: ", h.data)
	h.insert(40)
	fmt.Println("h.data: ", h.data)
	h.insert(150)
	fmt.Println("h.data: ", h.data)
	h.insert(35)
	fmt.Println("h.data: ", h.data)

	h.deleteMin()
	fmt.Println("h.data: ", h.data)
	h.deleteMin()
	fmt.Println("h.data: ", h.data)
	h.deleteMin()
	fmt.Println("h.data: ", h.data)
	h.deleteMin()
	fmt.Println("h.data: ", h.data)
	h.deleteMin()
	fmt.Println("h.data: ", h.data)
	h.deleteMin()
	fmt.Println("h.data: ", h.data)
}

func newHeap() heap {
	return heap{
		data:     []int{},
		quantity: 0,
	}
}

type heap struct {
	data     []int
	quantity int
}

func (h *heap) insert(number int) {
	if h.quantity < 0 {
		h.quantity = 0
	}

	h.data = append(h.data, number)
	h.bubbleUp(h.quantity)
	h.quantity++
}

func (h *heap) deleteMin() int {
	fmt.Println()
	if h.quantity == 0 {
		return 0
	}

	rightMost := h.quantity - 1
	min := h.data[0]
	h.data[0] = h.data[rightMost]

	h.quantity--
	h.data = h.data[:h.quantity]
	h.bubbleDown(0)

	return min
}

func (h *heap) bubbleDown(index int) {
	for {
		if index < 0 {
			break
		}

		left := index*2 + 1
		right := (index * 2) + 2
		if left >= h.quantity || right >= h.quantity {
			break
		}

		if h.data[index] > h.data[left] {
			h.exchange(index, left)
			index = left
		} else if h.data[index] > h.data[right] {
			h.exchange(index, right)
			index = right
		} else {
			break
		}
	}
}

func (h *heap) bubbleUp(index int) {
	for {
		if index <= 0 {
			break
		}

		parentIndex := (index / 2)
		if parentIndex < 0 {
			parentIndex = 0
		}

		if h.data[parentIndex] < h.data[index] {
			break
		}

		h.exchange(parentIndex, index)
		index = parentIndex
	}
}

func (h *heap) exchange(first, second int) {
	aux := h.data[first]
	h.data[first] = h.data[second]
	h.data[second] = aux

}
