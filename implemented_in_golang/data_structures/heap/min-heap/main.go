package main

import "fmt"

func main() {
	h := newHeap()
	// toInsert := []int{10, 8, 7, 5, 1}
	toInsert := []int{857, 149, 920, 468, 623, 117, 984, 537, 51, 160, 512, 271, 852, 372, 728, 160, 512, 363, 292, 838, 802, 459, 961, 837, 165, 203, 133, 518, 184, 733}

	for _, n := range toInsert {
		fmt.Println("insert ", n)
		h.insert(n)
		fmt.Println("h.data: ", h.data)
	}

	fmt.Println()
	for {
		fmt.Println()
		if len(h.data) == 0 {
			break
		}

		fmt.Println("extract: ", h.extractMin())
		fmt.Println("h.data: ", h.data)
	}
}

func newHeap() heap {
	return heap{
		data: []int{},
	}
}

type heap struct {
	data []int
}

func (h *heap) insert(number int) {
	h.data = append(h.data, number)
	h.bubbleUp(len(h.data) - 1)
}

func (h *heap) extractMin() int {
	if len(h.data) == 0 {
		return 0
	}

	rightMost := len(h.data) - 1

	min := h.data[0]
	h.data[0] = h.data[rightMost]
	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(0)
	return min
}

func (h *heap) bubbleDown(index int) {
	for {
		if index < 0 || index > len(h.data)-1 {
			break
		}

		left := index*2 + 1
		right := (index * 2) + 2

		if left < len(h.data) && right < len(h.data) {
			minChild := h.data[left]
			posToExchange := left
			if h.data[right] < h.data[left] {
				minChild = h.data[right]
				posToExchange = right
			}

			if minChild < h.data[index] {
				h.exchange(index, posToExchange)
				index = posToExchange
			} else {
				break
			}

		} else if left < len(h.data) && h.data[index] > h.data[left] {
			h.exchange(index, left)
			index = left
		} else if right < len(h.data) && h.data[index] > h.data[right] {
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

		parentIndex := ((index + 1) / 2) - 1
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
