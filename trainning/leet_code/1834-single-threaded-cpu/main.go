package main

import (
	"container/heap"
	"fmt"
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

func getOrder(tasks [][]int) []int {
	alltasks := make(AllTasks, len(tasks))
	i := 0

	for _, task := range tasks {
		alltasks[i] = &Task{
			index:          i,
			enqueueTime:    task[0],
			processingTime: task[1],
		}
		i++
	}

	heap.Init(&alltasks)

	availableTasks := AvailableTasks{}
	output := []int{}
	time := 1

	for alltasks.Len() > 0 || availableTasks.Len() > 0 {
		taskToProcess, newTime := getTaskToProcess(&alltasks, &availableTasks, time)
		if taskToProcess == nil {
			time++
			continue
		}

		time = newTime
		time += taskToProcess.processingTime
		output = append(output, taskToProcess.index)
	}

	return output
}

func getTaskToProcess(alltasks *AllTasks, availableTasks *AvailableTasks, time int) (*Task, int) {
	wasAdvanced := false
	i := -1
	for alltasks.Len() > 0 {
		i++

		if i == 0 && availableTasks.Len() == 0 && alltasks.Len() > 0 && (*alltasks)[0].enqueueTime > time {
			time = (*alltasks)[0].enqueueTime
			wasAdvanced = true
		}

		if (*alltasks)[0].enqueueTime > time {
			break
		}

		nextTask := heap.Pop(alltasks).(*Task)
		if availableTasks.Len() == 0 {
			heap.Init(availableTasks)
		}

		heap.Push(availableTasks, nextTask)
		if wasAdvanced {
			break
		}
	}

	if availableTasks.Len() > 0 {
		availableTask := heap.Pop(availableTasks).(*Task)
		return availableTask, time
	}

	return nil, time
}

type Task struct {
	index          int
	enqueueTime    int
	processingTime int
}

// A AllTasks implements heap.Interface and holds Tasks.
type AllTasks []*Task

func (pq AllTasks) Len() int { return len(pq) }

func (pq AllTasks) Less(i, j int) bool {
	if pq[i].enqueueTime == pq[j].enqueueTime {
		return pq[i].processingTime < pq[j].processingTime
	}

	return pq[i].enqueueTime < pq[j].enqueueTime
}

func (pq AllTasks) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *AllTasks) Push(x interface{}) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *AllTasks) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// /////////////////////////////////////////////////////////////////////////
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
