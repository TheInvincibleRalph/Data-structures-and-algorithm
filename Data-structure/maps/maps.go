package main

import (
	"fmt"
	"hash/fnv"
)

const bucketcount = 10

type Node struct {
	key      string
	value    int
	next     *Node
	previous *Node
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
			// if current.key == key {
			// 	current.value = value //updates the value if the key already exists.
			// 	return
			// }
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

func (ht *HashTable) delete(key string, value int) bool {
	index := hash(key)
	current := ht.buckets[index]

	for current != nil {
		if current.key == key && current.value == value {
			//if the node to be deleted is the head node
			if current.previous == nil {
				current = current.next
				ht.buckets[index] = current //this updates the head of the linked list, successfully deleting the fisrt element of the array
				if current.next != nil {
					current.next.previous = nil
				}
			} else {
				current.previous.next = current.next
				if current.next != nil {
					current.next.previous = current.previous
				}
			}
			//if the node to be deleted is not the last node

			return true //indicate successful deletion
		}
		current = current.next //move to the next node
	}

	return false //indicate that the key was not found
}

func (ht *HashTable) print() {
	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			fmt.Printf("%s: %d\n", current.key, current.value)
			current = current.next
		}
	}
}

func (ht *HashTable) printBucket(index int) {
	current := ht.buckets[index]
	for current != nil {
		fmt.Printf("%s: %d\n", current.key, current.value)
		current = current.next
	}
}

func main() {
	ht := &HashTable{}
	ht.insert("score", 50)
	ht.insert("point", 60)
	ht.insert("gp", 10)
	ht.insert("gp", 12)
	ht.insert("gp", 22)
	ht.insert("gup", 12)
	ht.insert("good", 15)
	ht.insert("gross", 124)
	ht.insert("gross", 198)
	ht.insert("gross", 154)
	ht.insert("gross", 196)
	ht.insert("gross", 120)
	ht.print()
	ht.printBucket(4)

	// fmt.Println(ht)
	ht.delete("gross", 124)
	value, found := ht.get("gross")
	if found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}
}
