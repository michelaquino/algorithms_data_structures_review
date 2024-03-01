package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const (
	resizeFactor     = 2
	overloadedFactor = 0.75
	initialSize      = 10
)

func main() {
	h := newHashTable()

	for i := 0; i < 20; i++ {
		h.insert(fmt.Sprintf("key %d", i), i)
	}
	h.print()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	///////// SEARCH
	value1, err := h.search("key 1")
	fmt.Println("value1", value1)
	fmt.Println("err", err)

	value2, err := h.search("key 199")
	fmt.Println("value2", value2)
	fmt.Println("err", err)

	///////// DELETE
	h.delete("key 1")
	value1, err = h.search("key 1")
	fmt.Println("value1", value1)
	fmt.Println("err", err)
}

func newHashTable() hashTable {
	return hashTable{
		data: make([]*element, initialSize, initialSize*resizeFactor),
	}
}

type element struct {
	key   string
	value int
	next  *element
}

type hashTable struct {
	size  int
	slots int
	data  []*element
}

func (h *hashTable) insert(key string, value int) {
	h.resize()

	newElement := element{
		key:   key,
		value: value,
	}

	index := h.indexToInsert(key)

	if h.data[index] != nil {
		newElement.next = h.data[index]
	}

	h.data[index] = &newElement
	h.size++
}

func (h *hashTable) delete(key string) {
	index := h.hash(key)
	if len(h.data) < index {
		return
	}

	node := h.data[index]
	for {
		if node == nil {
			return
		}

		if node.key == key {
			h.data[index] = node.next
			return
		}

		if node.next.key == key {
			node.next = node.next.next
			h.size--
			return
		}

		node = node.next
	}
}

func (h hashTable) search(key string) (int, error) {
	index := h.hash(key)
	if len(h.data) < index {
		return 0, errors.New("key not found")
	}

	node := h.data[index]
	for {
		if node == nil {
			return 0, errors.New("key not found")
		}

		if node.key == key {
			return node.value, nil
		}

		node = node.next
	}
}

func (h *hashTable) indexToInsert(key string) int {
	index := h.hash(key)

	// index inside quantity of itens
	if index <= len(h.data)-1 {
		return index
	}

	resizeSlice := make([]*element, index+1, cap(h.data))
	copy(resizeSlice, h.data)

	h.data = resizeSlice
	return index
}

func (h *hashTable) hash(key string) int {
	f := fnv.New32()
	f.Write([]byte(key))

	return int(f.Sum32()) % cap(h.data)
}

func (h *hashTable) resize() {
	if !h.isOverload() {
		return
	}

	rehashTable := hashTable{
		data: make([]*element, len(h.data), cap(h.data)*resizeFactor),
	}

	for _, element := range h.data {
		if element == nil {
			continue
		}

		toInsert := element
		for {
			if toInsert == nil {
				break
			}

			rehashTable.insert(toInsert.key, toInsert.value)
			toInsert = toInsert.next
		}
	}

	h.data = rehashTable.data
}

func (h hashTable) isOverload() bool {
	return float32(h.size)/float32(cap(h.data)) > overloadedFactor
}

func (h hashTable) print() {
	for i, slot := range h.data {
		fmt.Println()
		fmt.Println("slot ", i)

		if slot == nil {
			fmt.Println("nil")
			continue
		}

		toPrint := slot
		for {
			if toPrint == nil {
				fmt.Println("nil")
				break
			}

			fmt.Printf("%s:%d -> ", toPrint.key, toPrint.value)
			toPrint = toPrint.next
		}
	}
}
