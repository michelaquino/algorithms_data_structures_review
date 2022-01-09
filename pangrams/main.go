package main

import (
	"fmt"
	"strings"
)

func main() {
	pangram := "We promptly judged antique ivory buckles for the next prize"
	notPangram := "We promptly judged antique ivory buckles for the prize"

	fmt.Println(pangrams(pangram))
	fmt.Println(pangrams(notPangram))
}

func pangrams(s string) string {
	lettersCountMap := map[string]int{
		" ": 0,
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	for _, letter := range s {
		letterStr := strings.ToLower(string(letter))
		lettersCountMap[letterStr] += 1
	}

	isPangram := true
	for _, count := range lettersCountMap {
		if count == 0 {
			isPangram = false
			break
		}
	}

	if isPangram {
		return "pangram"
	}

	return "not pangram"
}
