package main

type Deque struct {
	items []int
}

// add an item to the front of the deque
func (d *Deque) addFront(item int) {
	d.items = append([]int{item}, d.items...) //
}

//OR

// func (d *Deque) addFront(item int) {
// 	d.items = append(d.items, 0) // add a new element to the end of the slice to make space for the new item
// 	copy(d.items[1:], d.items) // shift all elements to the right by one thereby duplicating the first element
// 	d.items[0] = item // add the new item to the first index
// }
