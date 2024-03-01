package main

import "fmt"

func main() {
	testCases := []struct {
		s             string
		k             int
		expectdOutput int
	}{
		{
			s:             "aaabb",
			k:             3,
			expectdOutput: 3,
		},
		{
			s:             "ababbc",
			k:             2,
			expectdOutput: 5,
		},
		{
			s:             "ababacb",
			k:             3,
			expectdOutput: 0,
		},
		{
			s:             "bbaaacbd",
			k:             3,
			expectdOutput: 3,
		},
		{
			s:             "aaabbb",
			k:             3,
			expectdOutput: 6,
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := longestSubstring(t.s, t.k)
		fmt.Println("output: ", output)
		if output != t.expectdOutput {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput, output))
		}
	}

}

func longestSubstring(str string, target int) int {
	return longestSubstringDivideConquer(str, 0, len(str), target)
}

func longestSubstringDivideConquer(str string, start, end, target int) int {
	if end-start <= 0 {
		return 0
	}

	charCountMap := map[rune]int{}
	for i := start; i < end; i++ {
		charCountMap[rune(str[i])]++
	}

	for i := start; i < end; i++ {
		char := rune(str[i])

		if charCountMap[char] >= target {
			continue
		}

		firstPartCount := longestSubstringDivideConquer(str, start, i, target)
		secondPartCount := longestSubstringDivideConquer(str, i+1, end, target)
		return max(firstPartCount, secondPartCount)
	}

	return end - start
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
