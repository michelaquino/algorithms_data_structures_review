package main

import "fmt"

func main() {

	head := &ListNode{
		Val: 1,
	}

	four := &ListNode{
		Val:  4,
		Next: head,
	}

	three := &ListNode{
		Val:  3,
		Next: four,
	}

	two := &ListNode{
		Val:  2,
		Next: three,
	}

	head.Next = two

	fmt.Printf("1 - %p: \n", head)
	fmt.Printf("2 - %p: \n", two)
	fmt.Printf("3 - %p: \n", three)
	fmt.Printf("4 - %p: \n", four)
	fmt.Println()

	fmt.Println(hasCycle(head))
}

/**
 * https://leetcode.com/problems/linked-list-cycle
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for {
		if &(*slow) == &(*fast) {
			return true
		}

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}
