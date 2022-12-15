package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/k-closest-points-to-origin/

func main() {

	// points := [][]int{{1, 3}, {-2, 2}}
	// k := 1
	points := [][]int{{6, 10}, {-3, 3}, {-2, 5}, {0, 2}}
	k := 3

	result := kClosest(points, k)
	fmt.Println("result: ", result)
}

func kClosest(points [][]int, k int) [][]int {

	minHeap := newHeap()
	for _, p := range points {
		minHeap.insert(point{x: p[0], y: p[1]})
	}

	result := [][]int{}
	for i := 0; i < k; i++ {
		result = append(result, minHeap.extractMin().arrayRepr())
	}

	return result
}

func newHeap() heap {
	return heap{
		data: []point{},
	}
}

type heap struct {
	data []point
}

func (h *heap) insert(p point) {
	h.data = append(h.data, p)
	h.bubbleUp(len(h.data) - 1)
}

func (h *heap) extractMin() point {
	if len(h.data) == 0 {
		return point{}
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
			if h.data[right].distanceToCenter() < h.data[left].distanceToCenter() {
				minChild = h.data[right]
				posToExchange = right
			}

			if minChild.distanceToCenter() < h.data[index].distanceToCenter() {
				h.exchange(index, posToExchange)
				index = posToExchange
			} else {
				break
			}

		} else if left < len(h.data) && h.data[index].distanceToCenter() > h.data[left].distanceToCenter() {
			h.exchange(index, left)
			index = left
		} else if right < len(h.data) && h.data[index].distanceToCenter() > h.data[right].distanceToCenter() {
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

		if h.data[parentIndex].distanceToCenter() < h.data[index].distanceToCenter() {
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

type point struct {
	x int
	y int
}

func (p point) distanceToCenter() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
}

func (p point) arrayRepr() []int {
	return []int{p.x, p.y}
}
