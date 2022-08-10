package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	hashTable := newHashTableStaticArray(20)
	hashTable.add("aasdasdasd", 1)
	hashTable.add("basdasqweqwe", 2)
	hashTable.add("cpqwepoqe", 3)
	hashTable.add("cpqwepoqe", 4)

	hashTable.print()
}

type hashTableStaticArray struct {
	table []*node
}

func newHashTableStaticArray(length int) hashTableStaticArray {
	return hashTableStaticArray{
		table: make([]*node, length),
	}
}

func (h *hashTableStaticArray) add(key string, value any) {
	hash := calculateHash(key)

	index := 0
	if len(h.table) > 0 {
		index = int(hash % len(h.table))
	}

	if h.table[index] == nil {
		h.table[index] = newNode(key, value)
		return
	}

	h.table[index].insert(key, value)
}

func (h hashTableStaticArray) print() {
	for _, node := range h.table {
		if node == nil {
			fmt.Println("nil")
			continue
		}

		node.print()
		fmt.Println("")
	}
}

type node struct {
	key   string
	value any
	next  *node
}

func newNode(key string, value any) *node {
	return &node{
		key:   key,
		value: value,
	}
}

func (n *node) insert(key string, value any) {
	if n.key == key {
		n.value = value
		return
	}

	if n.next == nil {
		next := newNode(key, value)
		n.next = next
		return
	}

	n.next.insert(key, value)
}

func (n node) print() {
	fmt.Printf("--> ")
	fmt.Printf("%v ", n.value)

	if n.next != nil {
		n.next.print()
	}
}

func calculateHash(key string) int {
	hash := fnv.New32()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}
