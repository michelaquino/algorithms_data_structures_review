package main

import "fmt"

// https://leetcode.com/problems/subtree-of-another-tree
func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	subRoot := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println("isSubtree: ", isSubtree(root, subRoot))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isIdentical(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isIdentical(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return (root == nil) == (subRoot == nil)
	}

	return root.Val == subRoot.Val && isIdentical(root.Left, subRoot.Left) && isIdentical(root.Right, subRoot.Right)
}
