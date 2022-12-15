package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		s1            string
		s2            string
		expectdOutput bool
	}{
		{
			s1:            "ab",
			s2:            "eidbaooo",
			expectdOutput: true,
		},
		{
			s1:            "ab",
			s2:            "eidboaoo",
			expectdOutput: false,
		},
		{
			s1:            "adc",
			s2:            "dcda",
			expectdOutput: true,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := checkInclusion(t.s1, t.s2)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func checkInclusion(s1 string, s2 string) bool {
	s1LetterCountMap := map[rune]int{}
	for _, letter := range s1 {
		s1LetterCountMap[letter]++
	}

	firstPointer := 0
	secondPointer := 0

	s2LetterCountMap := map[rune]int{}
	for secondPointer < len(s2) {
		s2LetterCountMap[rune(s2[secondPointer])]++
		if isPermutation(s1LetterCountMap, s2LetterCountMap) {
			return true
		}

		if secondPointer-firstPointer+1 < len(s1) {
			secondPointer++
			continue
		}

		if secondPointer-firstPointer+1 == len(s1) {
			s2LetterCountMap[rune(s2[firstPointer])]--
			firstPointer++
			secondPointer++
			continue
		}

		if secondPointer-firstPointer+1 > len(s1) {
			s2LetterCountMap[rune(s2[secondPointer])]--
			s2LetterCountMap[rune(s2[firstPointer])]--
			firstPointer++
			continue
		}
	}

	return false
}

func isPermutation(s1LetterCountMap, s2LetterCountMap map[rune]int) bool {
	for s1Letter, count := range s1LetterCountMap {
		s2LetterCount := s2LetterCountMap[s1Letter]
		if s2LetterCount != count {
			return false
		}
	}

	return true
}

func printMap(letterCountMap map[rune]int) {
	fmt.Println("----> MAP")
	for letter, count := range letterCountMap {
		fmt.Printf("%s -> %d\n", string(letter), count)
	}
}
