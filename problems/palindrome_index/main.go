package main

import (
	"fmt"
)

func main() {
	// s := "bcbc"
	// s := "aaababaaaa"
	s := "aaa"
	// s := "aaab"
	// fmt.Println(isPalindrome(s))
	fmt.Println(palindromeIndex(s))
}

func palindromeIndex(s string) int32 {
	len := len(s)

	i, j := 0, len-1
	for ; i < len; i, j = i+1, j-1 {
		if s[i:i+1] != s[j:j+1] {
			break
		}
	}

	if i > j {
		return -1
	}

	for a, b := i+1, j; a < j && b > i+1; a, b = a+1, b-1 {
		if s[a:a+1] != s[b:b+1] {
			return int32(j)
		}
	}

	return int32(i)
}
