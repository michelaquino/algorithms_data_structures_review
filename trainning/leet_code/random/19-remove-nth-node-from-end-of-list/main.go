package main

import (
	"fmt"
)

// here
func main() {
	testCases := []struct {
		n             int
		head          *ListNode
		expectdOutput *ListNode
	}{
		// {
		// 	n: 2,

		// 	head: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val: 3,
		// 				Next: &ListNode{
		// 					Val: 4,
		// 					Next: &ListNode{
		// 						Val: 5,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	expectdOutput: &ListNode{
		// 		Val: 3,
		// 	},
		// },
		// {
		// 	n: 1,
		// 	head: &ListNode{
		// 		Val: 1,
		// 	},
		// 	expectdOutput: nil,
		// },
		// {
		// 	n: 1,
		// 	head: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 		},
		// 	},
		// 	expectdOutput: &ListNode{
		// 		Val: 1,
		// 	},
		// },
		{
			n: 2,
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			expectdOutput: nil,
		},
		// {
		// 	n: 1,
		// 	head: &ListNode{
		// 		Val: 1,
		// 		Next: &ListNode{
		// 			Val: 2,
		// 			Next: &ListNode{
		// 				Val: 3,
		// 			},
		// 		},
		// 	},
		// 	expectdOutput: nil,
		// },
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := removeNthFromEnd(t.head, t.n)
		print(output)
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	if first == nil {
		return nil
	}

	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	if length == n {
		return head.Next
	}

	targetIndex := length - n - 1
	currentNode := head
	for i := 0; i < targetIndex; i++ {
		currentNode = currentNode.Next
	}

	currentNode.Next = currentNode.Next.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(head *ListNode) {
	fmt.Println("PRINT")
	next := head
	for next != nil {
		fmt.Printf("%d -> ", next.Val)
		next = next.Next
	}

	fmt.Println()
}
