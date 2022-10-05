package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/last-stone-weight/

func main() {
	stones := []int{377, 751, 685, 311, 232, 345, 819, 299, 858, 763, 913, 26, 348, 198, 837, 661, 221, 580, 940, 100, 260, 347, 780, 219, 133, 620, 291, 208, 857, 70}

	// fmt.Println("result: ", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	// fmt.Println("result: ", lastStoneWeight([]int{2, 2}))
	// fmt.Println("result: ", lastStoneWeight([]int{1}))
	// fmt.Println("result: ", lastStoneWeight([]int{5, 1, 8, 10, 7}))
	fmt.Println("result: ", lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	h := newHeap()
	for _, height := range stones {
		h.insert(height)
	}

	result := -1
	heaviestOne := 0
	heaviestTwo := 0
	for {
		sorted := make([]int, len(h.data))
		copy(sorted, h.data)
		sort.SliceStable(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		})

		if len(h.data) <= 1 {
			result = h.extractMax()
			break
		}

		heaviestOne = h.extractMax()
		heaviestTwo = h.extractMax()
		resultWeight := smash(heaviestOne, heaviestTwo)
		if resultWeight <= 0 {
			continue
		}

		h.insert(resultWeight)
	}

	return result
}

func smash(stone1, stone2 int) int {
	if stone1 == stone2 {
		return 0
	}

	if stone1 > stone2 {
		return stone1 - stone2
	}

	return stone2 - stone1
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

func (h *heap) extractMax() int {
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
			maxChild := h.data[left]
			posToExchange := left
			if h.data[right] > h.data[left] {
				maxChild = h.data[right]
				posToExchange = right
			}

			if h.data[index] < maxChild {
				h.exchange(index, posToExchange)
				index = posToExchange
			} else {
				break
			}

		} else if left < len(h.data) && h.data[index] < h.data[left] {
			h.exchange(index, left)
			index = left
		} else if right < len(h.data) && h.data[index] < h.data[right] {
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

		if h.data[parentIndex] > h.data[index] {
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
