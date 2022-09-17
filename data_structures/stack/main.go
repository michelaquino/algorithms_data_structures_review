package main

import (
	"errors"
	"fmt"
)

type stack struct {
	top  int
	data []int
}

func newStack() stack {
	return stack{
		top:  -1,
		data: make([]int, 1),
	}
}

func (s *stack) push(value int) {
	fmt.Println("top: ", s.top)
	fmt.Println("len(s.data): ", len(s.data))
	if s.top == len(s.data)-1 {
		fmt.Println(111)
		newData := make([]int, len(s.data)*2)
		copy(newData, s.data)
		s.data = newData
	}

	s.data[s.top+1] = value
	s.top++
}

func (s *stack) pop() (int, error) {
	if s.top == -1 {
		return -1, errors.New("underflow")
	}

	value := s.data[s.top]
	s.top--

	return value, nil
}

func main() {
	s := newStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	v, err := s.pop()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.pop()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.pop()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.pop()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.pop()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

}
