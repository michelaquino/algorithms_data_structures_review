package main

import "fmt"

func main() {
	s := newStringBuilder()
	s.write("a")
	s.write("b")
	s.write("c")
	s.write("d")

	fmt.Println(s.string())
}

type stringBuilder struct {
	stringArray []string
}

func newStringBuilder() stringBuilder {
	return stringBuilder{
		stringArray: make([]string, 1),
	}
}

func (s *stringBuilder) write(word string) {
	s.stringArray = append(s.stringArray, word)
}

func (s stringBuilder) string() string {
	result := ""

	for _, word := range s.stringArray {
		result += word
	}

	return result
}
