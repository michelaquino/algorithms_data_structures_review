package main

import "fmt"

func main() {
	l := linkedList[int]{}

	l.insert(10)
	l.insert(20)
	l.insert(30)
	l.insert(40)
	l.insert(50)
	l.print()

	l.delete(40)
	l.print()

	n := l.search(10)

	fmt.Println(n)
}

type linkedList[T comparable] struct {
	head *node[T]
}

func (l linkedList[T]) search(value T) *node[T] {
	return l.searchNode(l.head, value)
}

func (l linkedList[T]) searchNode(node *node[T], value T) *node[T] {
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	return l.searchNode(node.next, value)
}

func (l *linkedList[T]) insert(value T) {
	if l.head == nil {
		l.head = &node[T]{value: value}
		return
	}

	newHead := &node[T]{
		value: value,
		next:  l.head,
	}

	l.head = newHead
}

func (l *linkedList[T]) delete(value T) {
	l.deleteNode(l.head, value)
}

func (l *linkedList[T]) deleteNode(node *node[T], value T) {
	if node == nil || node.next == nil {
		return
	}

	if node.next.value == value {
		node.next = node.next.next
		return
	}

	l.deleteNode(node.next, value)
}

func (l linkedList[T]) print() {
	l.printList(l.head)
	fmt.Println()
}

func (l linkedList[T]) printList(node *node[T]) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.value)
	l.printList(node.next)
}

type node[T any] struct {
	value T
	next  *node[T]
}
