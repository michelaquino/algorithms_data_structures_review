package main

import "fmt"

func main() {
	findSubSets([]int{1, 2, 3})
}

func findSubSets(set []int) {
	backtracking(set, []int{}, 0)
}

func backtracking(set, subset []int, position int) {
	// is a solution?
	if position == len(set) {
		fmt.Println(subset)
		return
	}

	possibilities := []bool{true, false}
	for _, isTrue := range possibilities {
		if isTrue {
			backtracking(set, append(subset, set[position]), position+1)
			continue
		}

		backtracking(set, subset, position+1)
	}
}
