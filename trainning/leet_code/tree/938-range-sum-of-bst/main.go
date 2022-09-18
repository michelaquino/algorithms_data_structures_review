package main

// https://leetcode.com/problems/range-sum-of-bst/

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0

	sum += rangeSumBST(root.Left, low, high)

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	sum += rangeSumBST(root.Right, low, high)

	return sum
}
