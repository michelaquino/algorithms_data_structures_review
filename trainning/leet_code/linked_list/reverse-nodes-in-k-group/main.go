package main

import "fmt"

// https://leetcode.com/problems/reverse-nodes-in-k-group/

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	// printList(head)
	// fmt.Println()
	// fmt.Println(findLast(head, 5))

	printList(reverseKGroup(head, 1))
	// reverseKGroup(head, 2)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	index, last := findLast(head, k)
	if last == nil && index < k {
		return head
	}

	current := head
	var prev *ListNode = nil

	for {
		next := current.Next
		current.Next = prev
		prev = current

		if last != nil && next == *(&last) {
			head.Next = reverseKGroup(last, k)
			head = current
			break
		}

		if next == nil {
			head = prev
			break
		}

		current = next
	}

	return head
}

func findLast(head *ListNode, k int) (int, *ListNode) {
	if head == nil {
		return 0, head
	}

	index := 0
	last := head

	for {
		index++
		last = last.Next
		if index == k || last == nil {
			break
		}
	}

	if index < k {
		return index, nil
	}

	return index, last
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
