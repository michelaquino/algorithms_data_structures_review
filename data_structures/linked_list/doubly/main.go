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
	l.printReverse()

	l.delete(40)
	l.print()
	l.printReverse()

	n := l.search(10)
	fmt.Println(n)
}

type linkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
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
		newNode := &node[T]{value: value}
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode := &node[T]{
		value: value,
		prev:  nil,
		next:  l.head,
	}

	l.head.prev = newNode
	l.head = newNode
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
		node.next.prev = node
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

func (l linkedList[T]) printReverse() {
	l.printReverseList(l.tail)
	fmt.Println()
}

func (l linkedList[T]) printReverseList(node *node[T]) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.value)
	l.printReverseList(node.prev)
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}
