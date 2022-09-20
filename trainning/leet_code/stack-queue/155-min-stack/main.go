package main

import "fmt"

// https://leetcode.com/problems/min-stack/

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println("obj.GetMin(): ", obj.GetMin())
	obj.Pop()
	fmt.Println("obj.Top(): ", obj.Top())
	fmt.Println("obj.GetMin(): ", obj.GetMin())
}

type MinStack struct {
	data *node
}

func Constructor() MinStack {
	return MinStack{
		data: nil,
	}
}

func (this *MinStack) Push(val int) {
	newNode := &node{
		value: val,
	}

	if this.data == nil {
		newNode.minValue = val
		this.data = newNode
		return
	}

	newNode.minValue = this.data.minValue
	if val < this.data.minValue {
		newNode.minValue = val
	}

	newNode.next = this.data
	this.data = newNode
}

func (this *MinStack) Pop() {
	this.data = this.data.next
}

func (this *MinStack) Top() int {
	return this.data.value
}

func (this *MinStack) GetMin() int {
	return this.data.minValue
}

type node struct {
	value    int
	minValue int
	next     *node
}
