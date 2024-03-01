package main

import "fmt"

// https://leetcode.com/problems/merge-nodes-in-between-zeros/

func main() {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 0,
				},
			},
		},
	}

	printList(mergeNodes(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	nodeSum := head.Next
	sum := 0
	for {
		sum += nodeSum.Val
		nodeSum = nodeSum.Next

		if nodeSum == nil || nodeSum.Val == 0 {
			break
		}
	}

	newNode := &ListNode{Val: sum}
	newNode.Next = mergeNodes(nodeSum)
	return newNode
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.Val)
	printList(node.Next)
}
