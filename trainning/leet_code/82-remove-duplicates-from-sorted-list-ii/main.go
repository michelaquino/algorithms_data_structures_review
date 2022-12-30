package main

import (
	"fmt"
	"reflect"
)

// https://leetcode.com/problems/reverse-linked-list/

func main() {
	testCases := []struct {
		head          *ListNode
		expectdOutput *ListNode
	}{
		// {
		// 	head: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val: 3,
		// 				Next: &ListNode{
		// 					Val: 3,
		// 					Next: &ListNode{
		// 						Val: 4,
		// 						Next: &ListNode{
		// 							Val: 4,
		// 							Next: &ListNode{
		// 								Val:  5,
		// 								Next: nil,
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	expectdOutput: &ListNode{},
		// },
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			expectdOutput: &ListNode{},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := deleteDuplicates(t.head)
		fmt.Printf("output: '%v'\n", output)

		if !reflect.DeepEqual(output, t.expectdOutput) {
			panic(fmt.Sprintf("expcted '%v', got '%v'", t.expectdOutput, output))
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	sentinel := &ListNode{
		Val:  0,
		Next: head,
	}

	pred := sentinel
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}

			pred.Next = head.Next
		} else {
			pred = pred.Next
		}

		head = head.Next
	}

	return sentinel.Next
}

func deleteDuplicates_worse(head *ListNode) *ListNode {
	foundMap := map[int]bool{}
	var toReturn *ListNode
	var newNode *ListNode

	first := head
	for first != nil {
		if first.Next != nil && first.Val == first.Next.Val {
			foundMap[first.Val] = true
			first = first.Next
			continue
		}

		if _, exists := foundMap[first.Val]; exists {
			first = first.Next
			continue
		}

		if toReturn == nil {
			newNode = &ListNode{Val: first.Val}
			toReturn = newNode
		} else {
			newNode.Next = &ListNode{Val: first.Val}
			newNode = newNode.Next
		}

		first = first.Next
	}

	return toReturn
}

func print(head *ListNode) {
	fmt.Println()

	next := head
	for next != nil {
		fmt.Printf("%d -> ", next.Val)
		next = next.Next
	}
	fmt.Println()
}
