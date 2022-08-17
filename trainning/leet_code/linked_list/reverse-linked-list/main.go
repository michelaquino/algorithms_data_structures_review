package main

import (
	"fmt"
)

// https://leetcode.com/problems/reverse-linked-list/

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	printList(head)
	println()
	printList(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode = nil
	current := head

	for {
		next := current.Next
		current.Next = prev
		prev = current
		current = next

		if current == nil {
			break
		}
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.Val)
	printList(node.Next)
}
