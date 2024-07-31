package main

import "fmt"

type Deque struct {
	items []int
}

// add an item to the front of the deque
func (d *Deque) addFront(item int) {
	d.items = append([]int{item}, d.items...) //
}

//OR

func (d *Deque) anotherAddFront(item int) {
	d.items = append(d.items, 0) // add a new element to the end of the slice to make space for the new item
	copy(d.items[1:], d.items)   // shift all elements to the right by one thereby duplicating the first element
	d.items[0] = item            // add the new item to the first index
}

// add an item to the back of the deque
func (d *Deque) addBack(item int) {
	d.items = append(d.items, item)
}

// remove the first item from the deque
func (d *Deque) removeFront() int {
	if len(d.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

// remove the last item from the deque
func (d *Deque) removeBack() int {
	if len(d.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item
}
