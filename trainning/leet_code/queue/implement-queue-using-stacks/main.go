package main

import "fmt"

// https://leetcode.com/problems/implement-queue-using-stacks/
func main() {
	obj := Constructor()
	obj.Push(1)
	// obj.data.print()
	// obj.aux.print()
	obj.Push(2)
	// obj.data.print()
	// obj.aux.print()

	param_3 := obj.Peek()
	fmt.Println("param_3: ", param_3)

	param_2 := obj.Pop()
	fmt.Println("param_2: ", param_2)

	param_4 := obj.Empty()
	fmt.Println("param_4: ", param_4)
}

type MyQueue struct {
	data stack
	aux  stack
}

func Constructor() MyQueue {
	return MyQueue{
		data: newStack(),
		aux:  newStack(),
	}
}

func (q *MyQueue) Push(x int) {
	if q.Empty() {
		q.data.push(x)
		return
	}

	transferData(&q.data, &q.aux)
	q.aux.push(x)
	transferData(&q.aux, &q.data)
}

func (q *MyQueue) Pop() int {
	if q.Empty() {
		return 0
	}

	return q.data.pop()
}

func (q *MyQueue) Peek() int {
	if q.Empty() {
		return 0
	}

	return q.data.peek()
}

func (q *MyQueue) Empty() bool {
	return q.data.empty()
}

func transferData(from, to *stack) {
	for {
		if from.empty() {
			break
		}

		value := from.pop()
		to.push(value)
	}
}

type stack struct {
	data []int
	top  int
}

func newStack() stack {
	return stack{
		data: make([]int, 100),
		top:  -1,
	}
}

func (s *stack) push(x int) {
	s.top++
	s.data[s.top] = x
}

func (s *stack) pop() int {
	value := s.data[s.top]
	s.data[s.top] = 0

	s.top--
	return value
}

func (s *stack) peek() int {
	return s.data[s.top]
}

func (s *stack) empty() bool {
	return s.top == -1
}

func (s *stack) print() {
	fmt.Println()
	fmt.Println(s.data)
}
