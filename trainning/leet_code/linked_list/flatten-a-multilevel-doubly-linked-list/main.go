package main

// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
import "fmt"

func main() {
	// twelve := NewNode(12, nil, nil, nil)
	// eleven := NewNode(11, nil, twelve, nil)
	// ten := NewNode(10, nil, nil, nil)
	// nine := NewNode(9, nil, ten, nil)
	// eight := NewNode(8, nil, nine, eleven)
	// seven := NewNode(7, nil, eight, nil)
	// six := NewNode(6, nil, nil, nil)
	// five := NewNode(5, nil, six, nil)
	// four := NewNode(4, nil, five, nil)
	// three := NewNode(3, nil, four, seven)
	// two := NewNode(2, nil, three, nil)
	// one := NewNode(1, nil, two, nil)

	// twelve.Prev = eleven
	// eleven.Prev = nil
	// ten.Prev = nine
	// nine.Prev = eight
	// eight.Prev = seven
	// seven.Prev = nil
	// six.Prev = five
	// five.Prev = four
	// four.Prev = three
	// three.Prev = two
	// two.Prev = one
	// one.Prev = nil

	three := NewNode(3, nil, nil, nil)
	two := NewNode(2, nil, nil, nil)
	one := NewNode(1, nil, nil, nil)

	one.Child = two
	two.Child = three

	f := flatten(one)
	printList(f)
	fmt.Println()
	printReverseList(tail(f))
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	nextFlatened := flatten(root.Next)
	root.Next = nextFlatened
	if nextFlatened != nil {
		nextFlatened.Prev = root
	}

	childFlatened := flatten(root.Child)
	if childFlatened != nil {
		childTail := tail(childFlatened)

		if root.Next != nil {
			root.Next.Prev = childTail
		}

		childTail.Next = root.Next

		root.Child.Prev = root
		root.Next = root.Child
		root.Child = nil
	}

	return root
}

func tail(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}

	return tail(root.Next)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func NewNode(val int, prev, next, child *Node) *Node {
	return &Node{
		Val:   val,
		Prev:  prev,
		Next:  next,
		Child: child,
	}
}

func printList(node *Node) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.Val)
	printList(node.Next)
}

func printReverseList(node *Node) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.Val)
	printReverseList(node.Prev)
}
