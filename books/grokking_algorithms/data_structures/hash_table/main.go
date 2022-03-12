package main

import (
	"fmt"
)

const (
	defaultInitialSize = 5
	maxLoadFactor      = 0.7
)

func main() {
	hash := NewHashTable()

	hash.Add(1, "value-1")
	hash.Add(2, "value-2")
	hash.Add(3, "value-3")
	hash.Add(4, "value-4")
	hash.Add(5, "value-5")
	hash.Add(6, "value-6")

	hash.displayAll()
	hash.Del(3)
	hash.displayAll()

	fmt.Println("hash.Get(1): ", hash.Get(1))
	fmt.Println("hash.Get(2): ", hash.Get(2))
	fmt.Println("hash.Get(3): ", hash.Get(3))
	fmt.Println("hash.Get(4): ", hash.Get(4))
	fmt.Println("hash.Get(5): ", hash.Get(5))
	fmt.Println("hash.Get(6): ", hash.Get(6))
}

type HashTable struct {
	values         []string
	howManyRecords int
}

func NewHashTable() HashTable {
	values := make([]string, defaultInitialSize)

	return HashTable{
		values:         values,
		howManyRecords: 0,
	}
}

func (h *HashTable) Add(key int, value string) {
	position := h.hash(key)
	h.values[position] = value
	h.howManyRecords++

	h.redimention()
}

func (h *HashTable) Get(key int) string {
	position := h.hash(key)
	return h.values[position]
}

func (h *HashTable) Del(key int) {
	position := h.hash(key)
	h.values[position] = ""
	h.howManyRecords--
}

func (h HashTable) Size() int {
	return h.howManyRecords
}

func (h HashTable) displayAll() {
	fmt.Println("----------------")
	fmt.Println("How many records: ", h.howManyRecords)

	for index, value := range h.values {
		fmt.Printf("index: %d; value: %s\n", index, value)
	}
}

func (h *HashTable) redimention() {
	loadFactor := float32(h.howManyRecords) / float32(len(h.values))
	if loadFactor <= maxLoadFactor {
		return
	}

	newValues := make([]string, len(h.values)*2)
	copy(newValues, h.values)
	h.values = newValues
}

func (h *HashTable) hash(key int) int {
	records := len(h.values)
	if records == 0 {
		return 0
	}

	return key % records
}
