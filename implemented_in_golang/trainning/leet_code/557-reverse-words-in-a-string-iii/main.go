package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	testCases := []struct {
		str           string
		expectdOutput string
	}{
		{
			str:           "Let's take LeetCode contest",
			expectdOutput: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			str:           "God Ding",
			expectdOutput: "doG gniD",
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := reverseWords(t.str)
		fmt.Println("output: ", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted %v, got %v", t.expectdOutput, output))
		}
	}
}

func reverseWords(s string) string {
	wordList := strings.Split(s, " ")
	for i := 0; i < len(wordList); i++ {

		reversed := reverseString(wordList[i])
		wordList[i] = reversed
	}

	return strings.Join(wordList, " ")
}

func reverseString(word string) string {
	str := []byte(word)

	start := 0
	end := len(str) - 1

	for start < end {
		aux := str[end]
		str[end] = str[start]
		str[start] = aux

		start++
		end--
	}

	return string(str)
}
