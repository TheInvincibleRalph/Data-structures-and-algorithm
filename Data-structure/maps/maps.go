package main

import (
	"fmt"
	"hash/fnv"
)

const bucketcount = 10 //initial size of the array

type Node struct {
	key      string
	value    int
	next     *Node
	previous *Node
}

type HashTable struct {
	buckets []*Node //each element in the array is a pointer to a Node structure
	size    int
	count   int
}

func hash(key string, size int) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))                 //writes the key (converted to bytes)  to the hash function
	return int(hasher.Sum32() % uint32(size)) //this makes sure the value returned fits within the array bounds or size of the array
}

func (ht *HashTable) insert(key string, value int) {
	index := hash(key, ht.size)
	newNode := &Node{key: key, value: value}

	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		current := ht.buckets[index]
		for current.next != nil {
			if current.key == key {
				current.value = value //updates the value if the key already exists.
				return
			}
			current = current.next
		}
		current.next = newNode
		newNode.previous = current

	}
	ht.count++

	if float64(ht.count)/float64(ht.size) > 0.7 {
		ht.resize(ht.size * 2) //doubles the array size when the load factor is greater than 0.7
	}
}

func (ht *HashTable) get(key string) (int, bool) {
	index := hash(key, ht.size)
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
	index := hash(key, ht.size)
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
			ht.count--
			return true //indicate successful deletion
		}
		current = current.next //move to the next node
	}

	return false //indicate that the key was not found
}

// print function prints all the elements in the hash table
func (ht *HashTable) print() {
	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			fmt.Printf("%s: %d\n", current.key, current.value)
			ht.count++
			current = current.next
		}
	}
}

// printBucket function prints the elements in a bucket
func (ht *HashTable) printBucket(index int) {
	current := ht.buckets[index]
	for current != nil {
		fmt.Printf("%s: %d\n", current.key, current.value)
		current = current.next
	}
}

// printBucketCount function prints the number of elements in each bucket
func (ht *HashTable) printBucketCount() {
	for i, bucket := range ht.buckets {
		current := bucket
		fmt.Printf("Bucket %d -> ", i)
		for current != nil {
			fmt.Printf("%s: %d\n", current.key, current.value)
			current = current.next
		}
	}
}

// search function searches for a key in the hash table
func (ht *HashTable) search(key string) (int, bool) {
	index := hash(key, ht.size)
	current := ht.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// NewHashTable creates a new hash table with the specified size
func NewHashTable(s int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, s), //creates an array of pointers to Node structures
		size:    s,
	}
}

// resize function resizes the array when the load factor is greater than 0.7
func (ht *HashTable) resize(s int) {
	new := NewHashTable(s) //creates a new hash table with the new size
	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			new.insert(current.key, current.value) //inserts the key-value pair into the new hash table
			current = current.next

		}
	}
	ht.buckets = new.buckets //copies the new hash table to the old hash table
	ht.size = new.size       //updates the size of the old hash table
}

// loadFactor function calculates the load factor of the hash table
func (ht *HashTable) loadFactor() float64 {
	for _, bucket := range ht.buckets {
		current := bucket
		for current != nil {
			ht.count++
			current = current.next
		}
	}
	return float64(ht.count) / float64(ht.size)
}

func main() {
	ht := NewHashTable(bucketcount)
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
	ht.printBucketCount()
	ht.search("gross")
	ht.loadFactor()
	ht.resize(20)

	// fmt.Println(ht)
	ht.delete("gross", 124)
	value, found := ht.get("gross")
	if found {
		fmt.Println("Found value:", value)
	} else {
		fmt.Println("Value not found")
	}
}
