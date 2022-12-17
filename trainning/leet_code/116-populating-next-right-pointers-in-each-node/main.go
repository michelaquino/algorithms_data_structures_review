package main

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node
func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	link(root.Left, root.Right)
	connect(root.Left)
	connect(root.Right)
	return root
}

func link(left, right *Node) {
	if left == nil {
		return
	}

	left.Next = right
	linkEdgeSubtrees(left, right)
}

func linkEdgeSubtrees(node1, node2 *Node) {
	node1ToLink := node1.Right
	node2ToLink := node2.Left
	for node1ToLink != nil && node1ToLink.Next == nil {
		node1ToLink.Next = node2ToLink

		node1ToLink = node1ToLink.Right
		node2ToLink = node2ToLink.Left
	}
}
