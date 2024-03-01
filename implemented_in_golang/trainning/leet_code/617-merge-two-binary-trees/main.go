package main

// https://leetcode.com/problems/merge-two-binary-trees
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}

	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
