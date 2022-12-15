package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/single-threaded-cpu/

func main() {
	// tasks := [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}
	tasks := [][]int{{35, 36}, {11, 7}, {15, 47}, {34, 2}, {47, 19}, {16, 14}, {19, 8}, {7, 34}, {38, 15}, {16, 18}, {27, 22}, {7, 15}, {43, 2}, {10, 5}, {5, 4}, {3, 11}}

	// tasks := [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}
	// tasks := [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}
	// tasks := [][]int{{100, 100}, {1000000000, 1000000000}}
	// tasks := [][]int{{100, 100}, {200, 200}}
	output := getOrder(tasks)

	fmt.Println("output: ", output)
}

func getOrder(input [][]int) []int {
	tasks := make([]*Task, len(input))
	for i, task := range input {
		tasks[i] = &Task{
			index:          i,
			enqueueTime:    task[0],
			processingTime: task[1],
		}

		tasks = append(tasks)
	}

	sort.SliceStable(tasks, func(i, j int) bool {
		return tasks[i].enqueueTime < tasks[j].enqueueTime
	})

	output := []int{}
	availableTasks := AvailableTasks{}
	currentTime := 1
	for {

		// update current time if top task enqueue time is greater than it
		if len(tasks) > 0 && len(availableTasks) == 0 && tasks[0].enqueueTime > currentTime {
			currentTime = tasks[0].enqueueTime
		}

		// Get all tasks with enqueue time less than or equal current time
		for {
			if len(tasks) == 0 {
				break
			}

			if tasks[0].enqueueTime > currentTime {
				break
			}

			nextTask := tasks[0]
			if len(tasks) > 0 {
				tasks = tasks[1:]
			}

			if availableTasks.Len() == 0 {
				heap.Init(&availableTasks)
			}

			heap.Push(&availableTasks, nextTask)
		}

		if len(availableTasks) == 0 {
			break
		}

		topTask := heap.Pop(&availableTasks).(*Task)
		currentTime += topTask.processingTime
		output = append(output, topTask.index)
	}

	return output
}

type Task struct {
	index          int
	enqueueTime    int
	processingTime int
}

type AvailableTasks []*Task

func (pq AvailableTasks) Len() int { return len(pq) }

func (pq AvailableTasks) Less(i, j int) bool {
	if pq[i].processingTime == pq[j].processingTime {
		return pq[i].index < pq[j].index
	}

	return pq[i].processingTime < pq[j].processingTime
}

func (pq AvailableTasks) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *AvailableTasks) Push(x interface{}) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *AvailableTasks) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
