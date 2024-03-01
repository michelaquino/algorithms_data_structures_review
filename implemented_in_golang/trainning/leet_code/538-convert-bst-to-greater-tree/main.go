package main

import "fmt"

// https://leetcode.com/problems/convert-bst-to-greater-tree/

var (
	treeSlice = []any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}
)

func main() {
	tree := createTree(0)
	// inOrderTranversal(tree)
	// preOrderTranversal(tree)
	tree = convertBST(tree)
	fmt.Println()
	preOrderTranversal(tree)
}

func createTree(rootIndex int) *TreeNode {
	if len(treeSlice) < rootIndex || treeSlice[rootIndex] == nil {
		return nil
	}

	root := &TreeNode{
		Val:   treeSlice[rootIndex].(int),
		Left:  createTree((rootIndex * 2) + 1),
		Right: createTree((rootIndex * 2) + 2),
	}

	return root
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	updateMaxSum(root, 0)
	return root
}

func updateMaxSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	if root.Right != nil {
		sum = updateMaxSum(root.Right, sum)
	}

	sum += root.Val
	root.Val = sum

	if root.Left != nil {
		sum = updateMaxSum(root.Left, sum)
	}

	return sum
}

func (n *TreeNode) print() {
	if n == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(n.Val)
}

func inOrderTranversal(n *TreeNode) {
	if n == nil {
		return
	}

	inOrderTranversal(n.Left)
	fmt.Printf("%d -> ", n.Val)
	inOrderTranversal(n.Right)
}

func preOrderTranversal(n *TreeNode) {
	if n == nil {
		return
	}

	fmt.Printf("%d -> ", n.Val)
	preOrderTranversal(n.Left)
	preOrderTranversal(n.Right)
}
