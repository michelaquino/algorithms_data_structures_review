package main

import "fmt"

// https://leetcode.com/problems/validate-binary-search-tree
func main() {
	tree1 := &TreeNode{
		Val: 2,

		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println("tree1: ", isValidBST(tree1))

	tree2 := &TreeNode{
		Val: 5,

		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println("tree2: ", isValidBST(tree2))

	tree3 := &TreeNode{
		Val: 5,

		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("tree3: ", isValidBST(tree3))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	minNode := min(root)
	if minNode != nil && minNode.Val > root.Val {
		return false
	}

	maxNode := max(root)
	if maxNode != nil && maxNode.Val < root.Val {
		fmt.Println("2")
		return false
	}

	minRightSubTree := min(root.Right)
	if minRightSubTree != nil && minRightSubTree.Val <= root.Val {
		return false
	}

	maxLeftSubTree := max(root.Left)
	if maxLeftSubTree != nil && maxLeftSubTree.Val >= root.Val {
		return false
	}

	isLeftValid := isValidBST(root.Left)
	isRightValid := isValidBST(root.Right)

	return isLeftValid && isRightValid
}

func min(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return min(root.Left)
}

func max(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		return root
	}

	return max(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
