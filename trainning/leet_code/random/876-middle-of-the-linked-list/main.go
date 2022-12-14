package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		str           *ListNode
		expectdOutput *ListNode
	}{
		{
			str: &ListNode{
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
			},
			expectdOutput: &ListNode{
				Val: 3,
			},
		},
		{
			str: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
			expectdOutput: &ListNode{
				Val: 4,
			},
		},
	}

	for _, t := range testCases {
		fmt.Println("=========================================")
		output := middleNode(t.str)
		fmt.Println("output: ", output)

		if t.expectdOutput.Val != output.Val {
			panic(fmt.Sprintf("expcted %d, got %d", t.expectdOutput.Val, output.Val))
		}
	}
}

func middleNode(head *ListNode) *ListNode {
	first := head
	if first == nil {
		return nil
	}

	if first.Next == nil {
		return first
	}

	second := first.Next
	for {
		if second == nil {
			return first
		}

		if second.Next == nil {
			first = first.Next
			second = second.Next
			continue
		}

		first = first.Next
		second = second.Next.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
