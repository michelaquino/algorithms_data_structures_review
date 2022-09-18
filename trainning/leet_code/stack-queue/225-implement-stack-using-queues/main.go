package main

import "fmt"

// https://leetcode.com/problems/implement-stack-using-queues/
/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	fmt.Println("param_3: ", param_3)

	param_2 := obj.Pop()
	fmt.Println("param_2: ", param_2)

	param_4 := obj.Empty()
	fmt.Println("param_4: ", param_4)
}

type MyStack struct {
	head *node
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	newNode := node{
		value: x,
	}

	if !this.Empty() {
		newNode.next = this.head
	}

	this.head = &newNode
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}

	val := this.head.value
	this.head = this.head.next
	return val
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}

	return this.head.value
}

func (this *MyStack) Empty() bool {
	return this.head == nil
}

type node struct {
	value int
	next  *node
}
