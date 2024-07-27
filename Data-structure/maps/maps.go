package main

import (
	"fmt"
	"hash/fnv"
)

const bucketcount = 10

type Node struct {
	key   string
	value int
	next  *Node
}

type HashTable struct {
	buckets [bucketcount]*Node //each element in the array is a pointer to a Node structure
}

func hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))                //writes the key (converted to bytes)  to the hash function
	return int(hasher.Sum32() % bucketcount) //this makes sure the value returned fits within the array bounds
}

func (ht *HashTable) insert(key string, value int) {
	index := hash(key)
	newNode := &Node{key: key, value: value}

	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		current := ht.buckets[index]
		for current.next != nil {
			if current.key == key {
				current.value = value //updates the value if the key alrrady exists.
				return
			}
			current = current.next
		}
		current.next = newNode
	}
}

func (ht *HashTable) get(key string) (int, bool) {
	index := hash(key)

	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func main() {
	ht := &HashTable{}
	ht.insert("score", 50)
	ht.insert("point", 9)

	value, found := ht.get("score")
	if found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}
}
