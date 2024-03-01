package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/most-profit-assigning-work/

func main() {
	// difficulty := []int{4, 2, 6, 8, 10} // 100
	// profit := []int{20, 10, 30, 40, 50}
	// worker := []int{4, 5, 6, 7}

	// difficulty := []int{85, 47, 57} // 0
	// profit := []int{24, 66, 99}
	// worker := []int{40, 25, 25}

	difficulty := []int{13, 37, 58} // 190
	profit := []int{4, 90, 96}
	worker := []int{34, 73, 45}

	fmt.Println(maxProfitAssignment(difficulty, profit, worker))
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type job struct {
		difficulty int
		profit     int
	}

	jobs := []job{}
	for i := range difficulty {
		jobs = append(jobs, job{
			difficulty: difficulty[i],
			profit:     profit[i],
		})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].difficulty < jobs[j].difficulty
	})

	sum := 0
	jobProfit := make([]int, len(worker))
	for i, workerHability := range worker {
		for _, j := range jobs {

			if j.difficulty > workerHability {
				break
			}

			jobProfit[i] = max(jobProfit[i], j.profit)
		}

		sum += jobProfit[i]
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
