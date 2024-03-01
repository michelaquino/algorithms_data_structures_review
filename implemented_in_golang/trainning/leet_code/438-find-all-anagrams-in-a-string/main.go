package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string
func main() {
	testCases := []struct {
		s             string
		p             string
		expectdOutput []int
	}{
		{
			s:             "cbaebabacd",
			p:             "abc",
			expectdOutput: []int{0, 6},
		},
		{
			s:             "abab",
			p:             "ab",
			expectdOutput: []int{0, 1, 2},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := findAnagrams(t.s, t.p)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

func findAnagrams(s string, p string) []int {
	result := []int{}
	lenP := len(p) - 1
	if len(s) < len(p) {
		return result
	}

	letterPMap := map[byte]int{}
	for _, l := range p {
		letterPMap[byte(l)]++
	}

	start := 0
	end := lenP
	letterSMap := map[byte]int{}
	for i := start; i <= end; i++ {
		letterSMap[byte(s[i])]++
	}

	for {
		if isAnagram(letterSMap, letterPMap) {
			result = append(result, start)
		}

		// slide window
		start++
		end++
		if end-start > lenP || end >= len(s) {
			break
		}

		// diminui letra anterior
		sLetter := s[start-1]
		quantityS := letterSMap[sLetter]
		if quantityS > 0 {
			letterSMap[sLetter]--
		} else {
			delete(letterSMap, sLetter)
		}

		// adiciona proxima letra
		letterSMap[s[end]]++
	}

	return result
}

func isAnagram(a, b map[byte]int) bool {
	for letterS, quantityLetterP := range b {
		if quantityLetterS, exists := a[letterS]; !exists || quantityLetterS != quantityLetterP {
			return false
		}
	}

	return true
}

func findAnagrams_timeout(s string, p string) []int {
	result := []int{}
	lenP := len(p) - 1
	if len(s) < len(p) {
		return result
	}

	letterPMap := map[byte]int{}
	for _, l := range p {
		letterPMap[byte(l)]++
	}

	start := 0
	end := lenP
	for end-start == lenP && end < len(s) {
		letterSMap := map[byte]int{}
		isAnagram := true

		for i := start; i <= end; i++ {
			letter := s[i]
			if _, exists := letterPMap[letter]; !exists {
				start = i + 1
				end = start + lenP
				isAnagram = false
				break
			}

			letterSMap[letter]++
		}

		if !isAnagram {
			continue
		}

		isAnagram = true
		for letterS, quantityLetterP := range letterSMap {
			if quantityLetterS, exists := letterPMap[letterS]; !exists || quantityLetterS != quantityLetterP {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			result = append(result, start)
		}

		start++
		end++
	}

	return result
}

func printMap(m map[byte]int) {
	for key, value := range m {
		fmt.Printf("%s: %d\n", string(key), value)
	}

	fmt.Println()
}
