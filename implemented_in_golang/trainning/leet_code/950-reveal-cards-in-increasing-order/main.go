package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/reveal-cards-in-increasing-order/
func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}

	output := deckRevealedIncreasing(deck)
	fmt.Println("Output: ", output)
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool {
		return deck[i] < deck[j]
	})

	indexes := []int{}
	for i := 0; i < len(deck); i++ {
		indexes = append(indexes, i)
	}

	result := make([]int, len(deck))
	for _, card := range deck {
		index := indexes[0]
		indexes = indexes[1:]
		result[index] = card

		if len(indexes) > 0 {
			aux := indexes[0]
			indexes = indexes[1:]
			indexes = append(indexes, aux)
		}
	}

	return result
}
