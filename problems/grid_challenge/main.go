package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	grid := []string{"mpxz", "abcd", "wlmf"}

	// grid := []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}
	fmt.Println(gridChallenge(grid))
}

func gridChallenge(grid []string) string {
	gridSorted := []string{}

	for _, row := range grid {
		sorted := []string{}
		for _, letter := range row {
			sorted = append(sorted, string(letter))
		}

		sort.Strings(sorted)
		gridSorted = append(gridSorted, strings.Join(sorted, ""))
	}

	for index := range gridSorted[0] {
		arrayToCheck := []string{}
		for _, row := range gridSorted {
			arrayToCheck = append(arrayToCheck, row[index:index+1])
		}

		if !sort.StringsAreSorted(arrayToCheck) {
			return "NO"
		}
	}

	return "YES"
}
