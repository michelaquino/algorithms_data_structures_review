package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/maximum-performance-of-a-team/

func main() {
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	n := 6
	k := 2

	fmt.Println(maxPerformance(n, speed, efficiency, k))
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	efficiencyMinHeap := MinHeap{}
	heap.Init(&efficiencyMinHeap)

	indexesGot := map[int]uint8{}

	// ideia semelhante ao 630

	for i := 0; i < n; i++ {
		fmt.Println("==============================")
		actualSpeed := speed[i]
		actualEfficiency := efficiency[i]
		fmt.Println("i:", i)
		fmt.Println("actualSpeed: ", actualSpeed)
		fmt.Println("actualEfficiency: ", actualEfficiency)
		fmt.Println("efficiencyMinHeap.Len(): ", efficiencyMinHeap.Len())

		for j := 0; j < i; j++ {

		}

		fmt.Println(indexesGot)

		if efficiencyMinHeap.Len() == k {
			minEfficiency := efficiencyMinHeap[0]
			fmt.Println("---> minEfficiency.index: ", minEfficiency.index)
			fmt.Println("---> minEfficiency.value: ", minEfficiency.value)
			fmt.Println("---> actualEfficiency: ", actualEfficiency)
			fmt.Println("---> speed[minEfficiency.index]: ", speed[minEfficiency.index])

			if actualEfficiency < minEfficiency.value {
				fmt.Println("---> 11111")
				continue
			}

			if actualSpeed < speed[minEfficiency.index] {
				fmt.Println("---> 22222")
				continue
			}

			heap.Pop(&efficiencyMinHeap)

			fmt.Println(indexesGot)
			delete(indexesGot, minEfficiency.index)
			fmt.Println(indexesGot)
		}

		indexesGot[i] = 1
		item := &Item{value: actualEfficiency, index: i}
		heap.Push(&efficiencyMinHeap, item)
	}

	fmt.Println(indexesGot)

	sum := 0
	min := int(1e6)
	for index, _ := range indexesGot {
		sum += speed[index]
		if efficiency[index] < min {
			min = efficiency[index]
		}
	}

	return sum * min
}

type Item struct {
	value int
	index int
}

type MinHeap []*Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// func (pq MinHeap) print() {

// 	for _, item := range pq {
// 		fmt.Printf("%+v \n", item)
// 	}
// }
