package main

import "fmt"

var (
	mapNumberOptions = map[int][]string{
		0: {"+"},
		1: {""},
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
)

/*
2: {"a", "b", "c"},
3: {"d", "e", "f"},
4: {"g", "h", "i"},
*/

func main() {
	numberDigited := []int{2, 3, 4}
	printPossibleWords(numberDigited, 0, "")
}

func printPossibleWords(numberDigited []int, actualNumberIndex int, option string) {
	if actualNumberIndex >= len(numberDigited) {
		fmt.Printf("%s ", option)
		return
	}

	optionsLetters := mapNumberOptions[numberDigited[actualNumberIndex]]
	for _, letter := range optionsLetters {
		printPossibleWords(numberDigited, actualNumberIndex+1, fmt.Sprintf("%s%s", option, letter))
	}
}
