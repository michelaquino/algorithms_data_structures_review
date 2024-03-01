package main

import (
	"fmt"
)

func main() {
	operation := []string{"MyCircularQueue", "enQueue", "deQueue", "enQueue", "enQueue", "deQueue", "isFull", "isFull", "Front", "deQueue", "enQueue", "Front", "enQueue", "enQueue", "Rear", "Rear", "deQueue", "enQueue", "enQueue", "Rear", "Rear", "Front", "Rear", "Rear", "deQueue", "enQueue", "Rear", "deQueue", "Rear", "Rear", "Front", "Front", "enQueue", "enQueue", "Front", "enQueue", "enQueue", "enQueue", "Front", "isEmpty", "enQueue", "Rear", "enQueue", "Front", "enQueue", "enQueue", "Front", "enQueue", "deQueue", "deQueue", "enQueue", "deQueue", "Front", "enQueue", "Rear", "isEmpty", "Front", "enQueue", "Front", "deQueue", "enQueue", "enQueue", "deQueue", "deQueue", "Front", "Front", "deQueue", "isEmpty", "enQueue", "Rear", "Front", "enQueue", "isEmpty", "Front", "Front", "enQueue", "enQueue", "enQueue", "Rear", "Front", "Front", "enQueue", "isEmpty", "deQueue", "enQueue", "enQueue", "Rear", "deQueue", "Rear", "Front", "enQueue", "deQueue", "Rear", "Front", "Rear", "deQueue", "Rear", "Rear", "enQueue", "enQueue", "Rear", "enQueue"}
	data := []any{81, 69, nil, 92, 12, nil, nil, nil, nil, nil, 28, nil, 13, 45, nil, nil, nil, 24, 27, nil, nil, nil, nil, nil, nil, 88, nil, nil, nil, nil, nil, nil, 53, 39, nil, 28, 66, 17, nil, nil, 47, nil, 87, nil, 92, 94, nil, 59, nil, nil, 99, nil, nil, 84, nil, nil, nil, 52, nil, nil, 86, 30, nil, nil, nil, nil, nil, nil, 45, nil, nil, 83, nil, nil, nil, 22, 77, 23, nil, nil, nil, 14, nil, nil, 90, 57, nil, nil, nil, nil, 34, nil, nil, nil, nil, nil, nil, nil, 49, 59, nil, 71}

	expected := []any{nil, true, true, true, true, true, false, false, 12, true, true, 28, true, true, 45, 45, true, true, true, 27, 27, 13, 27, 27, true, true, 88, true, 88, 88, 24, 24, true, true, 24, true, true, true, 24, false, true, 47, true, 24, true, true, 24, true, true, true, true, true, 53, true, 84, false, 53, true, 53, true, true, true, true, true, 66, 66, true, false, true, 45, 17, true, false, 17, 17, true, true, true, 23, 17, 17, true, false, true, true, true, 57, true, 57, 87, true, true, 34, 92, 34, true, 34, 34, true, true, 59, true}

	checkResult := func(expected, result any, queue MyCircularQueue) {
		if result != expected {
			fmt.Println("=====================")
			queue.print()
			panic(fmt.Sprintf("expected: %v; got: %v", expected, result))
		}
	}

	obj := Constructor(data[0].(int))
	for i, op := range operation[1:] {
		index := i + 1
		fmt.Println()
		fmt.Println("index: ", index)
		fmt.Println("op: ", op)

		switch op {
		case "enQueue":
			value := data[index]
			fmt.Println("value: ", value)
			result := obj.EnQueue(value.(int))
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("enQueue: ", result)
			obj.print()
		case "deQueue":
			result := obj.DeQueue()
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("deQueue: ", result)
			obj.print()
		case "isEmpty":
			result := obj.IsEmpty()
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("isEmpty: ", result)
		case "isFull":
			result := obj.IsFull()
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("isFull: ", result)
		case "Front":
			result := obj.Front()
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("Front: ", result)
		case "Rear":
			result := obj.Rear()
			resultExpected := expected[index]
			checkResult(resultExpected, result, obj)

			fmt.Println("Rear: ", result)
		}
	}
}

func (this MyCircularQueue) print() {
	fmt.Println("data: ", this.data)
	fmt.Println("head: ", this.head)
	fmt.Println("tail: ", this.tail)
	fmt.Println("count: ", this.count)
}

type MyCircularQueue struct {
	head  int
	tail  int
	count int
	data  []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head: -1,
		tail: -1,
		data: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	defer func() {
		this.count++
	}()

	if this.IsEmpty() {
		this.data[0] = value
		this.head = 0
		this.tail = 0
		return true
	}

	position := this.tail + 1
	if this.tail == len(this.data)-1 {
		position = 0
	}

	this.data[position] = value
	this.tail = position
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.data[this.head] = 0

	newHead := this.head + 1
	if this.head == len(this.data)-1 {
		newHead = 0
	}

	this.head = newHead
	this.count--

	if this.IsEmpty() {
		this.head = -1
		this.tail = -1
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == len(this.data)
}
