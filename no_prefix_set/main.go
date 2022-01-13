package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"defgab", "abcde", "aabcde", "bbbbbbbbbb", "jabjjjad", "aab"}
	// words := []string{
	// 	"aab",
	// 	"aac",
	// 	"aacghgh",
	// 	"aabghgh",
	// }
	noPrefix(words)
}

func noPrefix(words []string) {
	prefixesMap := map[string]int{}
	wordsMap := map[string]int{}

	for _, word := range words {
		if _, ok := prefixesMap[word]; ok {
			fmt.Println("BAD SET")
			fmt.Println(word)
			return
		}

		prefix := strings.Builder{}
		for _, letter := range word {
			prefix.WriteRune(letter)
			prefixesMap[prefix.String()] = 1

			if _, ok := wordsMap[prefix.String()]; ok {
				fmt.Println("BAD SET")
				fmt.Println(word)
				return
			}
		}

		wordsMap[word] = 1
	}

	fmt.Println("GOOD SET")
}
