package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

const (
	resizeFactor     = 2
	overloadedFactor = 0.7
	initialSize      = 2
)

func main() {
	h := newHashTable()

	for i := 0; i < 10; i++ {
		h.insert(fmt.Sprintf("key %d", i), i)

	}
	h.print()

	// ///////// SEARCH
	fmt.Println()
	value1, err := h.search("key 1")
	fmt.Println("value1", value1)
	fmt.Println("err", err)

	value2, err := h.search("not found")
	fmt.Println("value2", value2)
	fmt.Println("err", err)

	// ///////// DELETE
	fmt.Println()
	h.delete("key 1")

	fmt.Println()
	value1Deleted, err := h.search("key 1")
	fmt.Println("value1 deleted", value1Deleted)
	fmt.Println("err", err)

	fmt.Println()
	h.print()
}

func newHashTable() hashTable {
	return hashTable{
		data: make([]*element, initialSize, initialSize*resizeFactor),
	}
}

type element struct {
	key   string
	value int
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

	h.data[index] = &newElement
	h.size++
}

func (h *hashTable) delete(key string) {
	index := h.searchIndex(key)
	if index == -1 {
		return
	}

	for ; index < len(h.data); index++ {
		if index+1 == len(h.data) {
			h.data[index] = nil
			break
		}

		h.data[index] = h.data[index+1]
	}
}

func (h hashTable) search(key string) (*element, error) {
	index := h.hash(key)
	if len(h.data) < index {
		return nil, errors.New("key not found")
	}

	indexFound := h.searchIndex(key)
	if indexFound == -1 {
		return nil, errors.New("key not found")
	}

	return h.data[index], nil
}

func (h hashTable) searchIndex(key string) int {
	index := h.hash(key)
	if len(h.data) < index {
		return -1
	}

	for ; index < len(h.data); index++ {
		node := h.data[index]

		if node == nil {
			return -1
		}

		if node.key == key {
			return index
		}
	}

	return -1
}

func (h *hashTable) indexToInsert(key string) int {
	index := h.hash(key)

	defer func() {
		// index outside quantity of itens
		if index > len(h.data)-1 {
			resizeSlice := make([]*element, index+1)
			copy(resizeSlice, h.data)
			h.data = resizeSlice
		}
	}()

	// find another place to insert
	for ; index < len(h.data); index++ {
		if h.data[index] == nil {
			return index
		}
	}

	return index
}

func (h *hashTable) hash(key string) int {
	f := fnv.New32()
	f.Write([]byte(key))

	index := int(f.Sum32()) % cap(h.data)
	return index
}

func (h *hashTable) resize() {
	if !h.isOverload() {
		return
	}

	rehashTable := hashTable{
		data: make([]*element, cap(h.data)*resizeFactor),
	}

	for _, element := range h.data {
		if element == nil {
			continue
		}

		toInsert := element
		rehashTable.insert(toInsert.key, toInsert.value)
	}

	h.data = rehashTable.data
}

func (h hashTable) isOverload() bool {
	return float32(h.size)/float32(cap(h.data)) > overloadedFactor
}

func (h hashTable) print() {
	fmt.Println("===============")

	for i, slot := range h.data {
		toPrint := ""
		if slot == nil {
			toPrint = "nil"
		} else {
			toPrint = fmt.Sprintf("%s: %d ", slot.key, slot.value)
		}

		fmt.Printf("slot %d => %s\n", i, toPrint)
	}

	fmt.Println()
}
