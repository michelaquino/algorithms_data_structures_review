package main

import (
	"fmt"
)

type item struct {
	name   string
	weight int
	value  int
}

func main() {
	// itens := map[int]item{
	// 	1: {
	// 		name:   "item-1",
	// 		weight: 13,
	// 		value:  30000,
	// 	},
	// 	2: {
	// 		name:   "item-2",
	// 		weight: 5,
	// 		value:  9000,
	// 	},
	// 	3: {
	// 		name:   "item-3",
	// 		weight: 2,
	// 		value:  1000,
	// 	},
	// 	4: {
	// 		name:   "item-4",
	// 		weight: 4,
	// 		value:  2000,
	// 	},
	// 	5: {
	// 		name:   "item-5",
	// 		weight: 1,
	// 		value:  5000,
	// 	},
	// }

	// maxWeight := 10
	itens := map[int]item{
		0: {
			name:   "violão",
			weight: 1,
			value:  1500,
		},
		1: {
			name:   "rádio",
			weight: 4,
			value:  3000,
		},
		2: {
			name:   "notebook",
			weight: 3,
			value:  2000,
		},
	}

	maxWeight := 4

	stoleItens := findItensToStole(itens, maxWeight)
	fmt.Println(stoleItens)
}

func findItensToStole(itens map[int]item, maxWeight int) []item {
	// Matrix to store itens in each estimative
	stoleItensMatrix := [][][]item{}

	table := make([][]int, len(itens))
	stoleItensMatrix = make([][][]item, len(itens))
	for i := 0; i < len(itens); i++ {
		table[i] = make([]int, maxWeight)
		stoleItensMatrix[i] = make([][]item, maxWeight)
	}

	for i := 0; i < len(table); i++ {
		itemToCheck := itens[i]
		for j := 0; j < len(table[i]); j++ {

			// Does not fit in the sub-backpack
			if itemToCheck.weight > j+1 {

				// If it's not first line, get the previous line data
				if i != 0 {
					table[i][j] = table[i-1][j]
					stoleItensMatrix[i][j] = stoleItensMatrix[i-1][j]
				}

				continue
			}

			value := itemToCheck.value
			itensEstimative := []item{itemToCheck}

			// If it's not first line, get the previous line data
			if i > 0 {
				lastEstimative := table[i-1][j]
				actualEstimative := itemToCheck.value

				/*
					If it's not first column, the actual estimative
					is the actual item plus the item that fit the rest of space
				*/
				if j-itemToCheck.weight+1 > 0 {
					actualEstimative = itemToCheck.value + table[i][j-itemToCheck.weight+1]
				}

				if lastEstimative > actualEstimative {
					value = lastEstimative
					itensEstimative = stoleItensMatrix[i-1][j]
				} else {
					value = actualEstimative
					itensEstimative = append(stoleItensMatrix[i][j-itemToCheck.weight+1], itemToCheck)
				}
			}

			table[i][j] = value
			stoleItensMatrix[i][j] = itensEstimative
		}
	}

	// printMatrix(table)

	stoleItens := stoleItensMatrix[len(stoleItensMatrix)-1][len(stoleItensMatrix)-1]
	return stoleItens
}

func printMatrix(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%d ", table[i][j])
		}

		fmt.Println()
	}
}
