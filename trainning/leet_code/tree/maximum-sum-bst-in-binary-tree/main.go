package main

import (
	"fmt"
)

func createTree(treeSlice []any, rootIndex int) *TreeNode {
	if len(treeSlice)-1 < rootIndex || treeSlice[rootIndex] == nil {
		return nil
	}

	root := &TreeNode{
		Val:   treeSlice[rootIndex].(int),
		Left:  createTree(treeSlice, (rootIndex*2)+1),
		Right: createTree(treeSlice, (rootIndex*2)+2),
	}

	return root
}

// https://leetcode.com/problems/maximum-sum-bst-in-binary-tree/
func main() {
	slice := []any{4, 8, nil, 6, 1, 9, nil, -5, 4, nil, nil, nil, -3, nil, 10}
	// slice := []any{4, 3, nil, 1, 2}
	// slice := []any{-4, -2, -5}

	tree := createTree(slice, 0)

	max := maxSumBST(tree)

	fmt.Println(max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, _, _, max := solve(root)
	if max < 0 {
		return 0
	}
	return max
}

func solve(root *TreeNode) (min int, max int, isBst bool, sum int) {
	min, max = root.Val, root.Val

	leftSum := 0
	rightSum := 0

	actualIsBst := true
	if root.Left != nil {
		minLeft := 0
		maxLeft := 0
		leftIsBst := true

		minLeft, maxLeft, leftIsBst, leftSum = solve(root.Left)
		if minLeft > root.Val || maxLeft > root.Val || !leftIsBst {
			actualIsBst = false
		}

		min = minInt(min, minLeft)
		max = maxInt(max, maxLeft)
	}

	if root.Right != nil {
		minRight := 0
		maxRight := 0
		rightIsBst := true

		minRight, maxRight, rightIsBst, rightSum = solve(root.Right)
		if minRight < root.Val || maxRight < root.Val || !rightIsBst {
			actualIsBst = false
		}

		min = minInt(min, minRight)
		max = maxInt(max, maxRight)
	}

	if actualIsBst {
		sum = root.Val + leftSum + rightSum
	} else {
		sum = maxInt(leftSum, rightSum)
	}

	return min, max, actualIsBst, sum
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
