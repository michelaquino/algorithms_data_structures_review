package main

import "fmt"

// https://leetcode.com/problems/merge-k-sorted-lists/

func main() {
	// lists = [[1,4,5],[1,3,4],[2,6]]

	one := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	two := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	three := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}

	lists := []*ListNode{one, two, three}
	printList(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var min *ListNode = nil
	minIndex := -1

	for i, l := range lists {
		if l == nil {
			continue
		}

		if min == nil {
			min = l
			minIndex = i
			continue
		}

		if l.Val <= min.Val {
			min = l
			minIndex = i
		}
	}

	if min != nil {
		lists[minIndex] = lists[minIndex].Next
		min.Next = mergeKLists(lists)
	}

	return min
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printKLists(lists []*ListNode) {
	for _, l := range lists {
		printList(l)
		fmt.Println()
	}
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Printf("nil")
		return
	}

	fmt.Printf("%v --> ", node.Val)
	printList(node.Next)
}
