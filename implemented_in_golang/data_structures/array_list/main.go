package main

import "fmt"

const (
	resizeFactor = 2
)

func main() {
	array := newArrayList()
	array.append(1)
	array.append(2)
	array.append(3)
	array.append(4)

	array.print()
}

type arrayList struct {
	nextIndex int
	array     []any
}

func newArrayList() arrayList {
	return arrayList{
		nextIndex: 0,
		array:     make([]any, 1),
	}
}

func (a arrayList) print() {
	for _, v := range a.array {
		fmt.Printf("%v ", v)
	}
}

func (a *arrayList) append(value any) {
	a.resize()
	a.array[a.nextIndex] = value
	a.nextIndex++
}

func (a *arrayList) resize() {
	if a.nextIndex <= len(a.array)-1 {
		return
	}

	newArray := make([]any, len(a.array)*resizeFactor)
	copy(newArray, a.array)
	a.array = newArray
}
