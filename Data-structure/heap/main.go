package main

import "fmt"

type Heap struct {
	items []int
}

// <=================================Helper Functions=================================>

// returns the index of the parent node
func parent(i int) int {
	return (i - 1) / 2
}

// returns the index of the left child
func left(i int) int {
	return 2*i + 1
}

// returns the index of the right child
func right(i int) int {
	return 2*i + 2
}

// swaps two elements in the heap
func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// <=================================Heapify Functions=================================>

func (h *Heap) heapifyDown(i int) {
	for {
		left := left(i)
		right := right(i)
		smallest := i //initialize smallest to the parent

		//check if the left child is less than the parent
		if left < len(h.items) && h.items[left] < h.items[smallest] {
			smallest = left
		}

		//check if the right child is less than the parent
		if right < len(h.items) && h.items[right] < h.items[smallest] {
			smallest = right
		}

		//if the parent is less than the children, then the heap is valid
		if i == smallest {
			break
		}

		//swap the parent with the smallest child
		h.swap(i, smallest)
		i = smallest

	}
}

func (h *Heap) heapifyUp(i int) {
	for i > 0 && h.items[parent(i)] > h.items[i] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func main() {
	heap := Heap{}
	heap.items = []int{10, 5, 3, 2, 4}
	// heap.heapifyUp(len(heap.items) - 1) //heapify the heap starting from the last node
	// fmt.Println(heap.items)
	heap.heapifyDown(0) //heapify the heap starting from the root node (index 0)
	fmt.Println(heap.items)
}
