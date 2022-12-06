package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/binary-tree-paths/

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := binaryTreePaths(node)
	fmt.Println("result: ", result)

}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	answer := []string{}
	var constructStr func(root *TreeNode, partial []string)

	constructStr = func(root *TreeNode, partial []string) {
		if root == nil {
			return
		}

		actual := strconv.Itoa(root.Val)
		partial = append(partial, actual)
		if root.Left == nil && root.Right == nil {
			strResult := strings.Join(partial, "->")
			answer = append(answer, strResult)
			return
		}

		constructStr(root.Left, partial)
		constructStr(root.Right, partial)
	}

	constructStr(root, []string{})
	return answer
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
