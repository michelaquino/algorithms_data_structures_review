package main

import "fmt"

func main() {
	set := []int{1, 2, 3}

	// i := 2
	// fmt.Println("a: ", append(set[:i], set[i+1:]...))

	findSubSets(set)
}

func findSubSets(set []int) {
	backtracking(set, []int{}, 0)
}

func backtracking(set, subset []int, position int) {
	if position >= len(set)+len(subset) {
		fmt.Println(subset)
		return
	}

	for i, v := range set {
		setCopied := make([]int, len(set))
		copy(setCopied, set)

		newSet := append(setCopied[:i], setCopied[i+1:]...)
		backtracking(newSet, append(subset, v), position+1)
	}
}
