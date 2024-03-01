package main

import (
	"errors"
	"fmt"
)

func main() {
	s := newQueue()
	s.enqueue(1)
	s.enqueue(2)
	s.enqueue(3)
	s.enqueue(4)

	v, err := s.dequeue()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.dequeue()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.dequeue()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.dequeue()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)

	v, err = s.dequeue()
	fmt.Println("value: ", v)
	fmt.Println("error: ", err)
}

type queue struct {
	head *node
	tail *node
}

type node struct {
	value int
	prev  *node
	next  *node
}

func newQueue() queue {
	return queue{}
}

func (q *queue) enqueue(value int) {
	node := &node{
		value: value,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}

	q.head.prev = node
	node.next = q.head
	q.head = node
}

func (q *queue) dequeue() (value int, err error) {
	if q.head == nil || q.tail == nil {
		return 0, errors.New("empty")
	}

	value = q.tail.value
	if q.tail.prev == nil {
		q.head = nil
		q.tail = nil
		return
	}

	q.tail = q.tail.prev
	q.tail.next = nil
	return
}

func (q queue) print() {
	fmt.Println()
	fmt.Printf("head: %+v\n", q.head)
	fmt.Printf("tail: %+v\n", q.tail)

	toPrint := q.head
	for {
		if toPrint == nil {
			return
		}

		fmt.Printf("%+v -> ", toPrint)
		toPrint = toPrint.next
	}
}
