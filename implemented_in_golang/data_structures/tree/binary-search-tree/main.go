package main

import "fmt"

func main() {
	fmt.Printf("\n===================\n")
	fmt.Println("Insert")
	fmt.Println()
	root := &node{value: 5}
	root.insert(2)
	root.insert(1)
	root.insert(3)
	root.insert(7)
	root.insert(8)
	root.insert(6)

	// fmt.Printf("\n===================\n")
	// fmt.Println("Tranversal")
	// fmt.Println()
	// inOrderTranversal(root)
	// fmt.Println()
	// preOrderTranversal(root)
	// fmt.Println()
	// posOrderTranversal(root)

	// fmt.Printf("\n===================\n")
	// fmt.Println("Search")
	// fmt.Println()
	// n := root.search(5)
	// n.print()
	// n = root.search(2)
	// n.print()
	// n = root.search(1)
	// n.print()
	// n = root.search(3)
	// n.print()
	// n = root.search(6)
	// n.print()
	// n = root.search(7)
	// n.print()

	// fmt.Printf("\n===================\n")
	// fmt.Println("min")
	// fmt.Println()
	// min := root.min()
	// min.print()

	// fmt.Printf("\n===================\n")
	// fmt.Println("max")
	// fmt.Println()
	// max := root.max()
	// max.print()

	// fmt.Printf("\n===================\n")
	// fmt.Println("predecessor")
	// fmt.Println()
	// predecessor := root.predecessor(5)
	// predecessor.print()

	// fmt.Printf("\n===================\n")
	// fmt.Println("sucessor")
	// fmt.Println()
	// sucessor := root.sucessor(5)
	// sucessor.print()

	fmt.Printf("\n===================\n")
	fmt.Println("delete")
	fmt.Println()
	deleted := root.delete(7)
	deleted.print()

	fmt.Println()
	inOrderTranversal(root)
}

func inOrderTranversal(n *node) {
	if n == nil {
		return
	}

	inOrderTranversal(n.left)
	fmt.Printf("%d -> ", n.value)
	inOrderTranversal(n.right)
}

func preOrderTranversal(n *node) {
	if n == nil {
		return
	}

	fmt.Printf("%d -> ", n.value)
	preOrderTranversal(n.left)
	preOrderTranversal(n.right)
}

func posOrderTranversal(n *node) {
	if n == nil {
		return
	}

	posOrderTranversal(n.left)
	posOrderTranversal(n.right)
	fmt.Printf("%d -> ", n.value)
}

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) print() {
	if n == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(n.value)
}

func (n *node) insert(value int) {
	if n == nil {
		return
	}

	if value <= n.value {
		if n.left == nil {
			newNode := &node{value: value}
			n.left = newNode
			return
		}

		n.left.insert(value)
		return
	}

	if n.right == nil {
		newNode := &node{value: value}
		n.right = newNode
		return
	}

	n.right.insert(value)
}

func (n *node) delete(value int) *node {
	if n == nil {
		return nil
	}

	if value < n.value {
		if n.left == nil {
			return nil
		}

		if n.left.value == value {
			toReturn := n.left

			if n.left.left == nil && n.left.right == nil {
				n.left = nil
			} else if n.left.left == nil && n.left.right != nil {
				n.left = n.left.right
			} else if n.left.left != nil && n.left.right == nil {
				n.left = n.left.left
			} else {
				successor := n.left.sucessor(value)
				n.delete(successor.value)
				n.left.value = successor.value
			}

			return toReturn
		}

		deleted := n.left.delete(value)
		return deleted
	}

	if value > n.value {
		if n.right == nil {
			return nil
		}

		if n.right.value == value {
			toReturn := n.right

			if n.right.left == nil && n.right.right == nil {
				n.right = nil
			} else if n.right.left == nil && n.right.right != nil {
				n.right = n.right.right
			} else if n.right.left != nil && n.right.right == nil {
				n.right = n.right.left
			} else {
				successor := n.right.sucessor(value)
				n.delete(successor.value)
				n.right.value = successor.value
			}

			return toReturn
		}

		deleted := n.right.delete(value)
		return deleted
	}

	return nil
}

func (n *node) search(value int) *node {
	if n == nil {
		return nil
	}

	if n.value == value {
		return n
	}

	if value <= n.value {
		return n.left.search(value)
	}

	return n.right.search(value)
}

func (n *node) min() *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return n
	}

	return n.left.min()
}

func (n *node) max() *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return n
	}

	return n.right.min()
}

func (n *node) predecessor(value int) *node {
	if n == nil {
		return nil
	}

	if n.left == nil {
		return nil
	}

	return n.left.max()
}

func (n *node) sucessor(value int) *node {
	if n == nil {
		return nil
	}

	if n.right == nil {
		return nil
	}

	return n.right.min()
}
