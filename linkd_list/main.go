package main

import "fmt"

func main() {
	ll := LinkedList{}
	for i := 0; i < 10; i++ {
		newNode := Node{
			data: i,
		}

		ll.Add(&newNode)
	}

	ll.Print()
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(node *Node) {
	if node == nil {
		return
	}

	if l.head == nil {
		l.head = node
		return
	}

	l.head.Add(node)
}

func (l LinkedList) Print() {
	if l.head == nil {
		fmt.Println("Empty list")
		return
	}

	l.head.Print()
}

type Node struct {
	data int
	next *Node
}

func (n *Node) Add(next *Node) {
	if next == nil {
		return
	}

	if n.next == nil {
		n.next = next
		return
	}

	n.next.Add(next)
}

func (n Node) Print() {
	fmt.Println("data: ", n.data)
	if n.next == nil {
		return
	}

	n.next.Print()
}
