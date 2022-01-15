package main

import (
	"fmt"
	"math"
)

func toTest() {
	genes := []string{"a", "b", "c", "aa", "d", "b"}
	geneHealth := []int{1, 2, 3, 4, 5, 6}

	// d := "caaab"
	// return
	// first := 1
	// last := 5
	// d := "xyz"
	// first := 0
	// last := 4
	d := "bcdybc"
	first := 2
	last := 4

	// return

	fmt.Println(_dnaHealth(genes, geneHealth, d, first, last))
}

func _dnaHealth(genes []string, geneHealth []int, d string, first, last int) int64 {
	min := 2 * int64(math.Pow(10, 6))
	max := -1
	fmt.Println(min)
	fmt.Println(max)
	return 0

	// var sum int64 = 0
	// for i := first; i <= last; i++ {
	// 	gene := genes[i]

	// 	count := _naiveSearch(d, gene)

	// 	geneHealthSum := int64(count * geneHealth[i])

	// 	if geneHealthSum < min {
	// 		min = geneHealthSum
	// 	}

	// 	if geneHealthSum > max {
	// 		max = geneHealthSum
	// 	}

	// 	sum += geneHealthSum
	// }

	// return sum
}

func _naiveSearch(str, pattern string) int {
	lenPattern := len(pattern)
	j := 0
	quantityFound := 0

	startSearchIndex := 0
	for i := 0; i < len(str); i++ {

		if string(str[i]) != string(pattern[j]) {
			startSearchIndex++
			i = startSearchIndex - 1
			j = 0
			continue
		}

		j++
		// Not found
		if j < lenPattern {
			continue
		}

		quantityFound++
		j = 0
		startSearchIndex = i

		if i+1 >= len(str) || i == 0 {
			continue
		}

		/*
			Ex: pattern: aa -> str: aaa.
			Não pode incrementar 'i', pq nesse exemplo o 2o 'a' é final e inicio de um pattern
		*/
		if len(pattern) > 1 {
			i--
		}
	}

	return quantityFound
}
