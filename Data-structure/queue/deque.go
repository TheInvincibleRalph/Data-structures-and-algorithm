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

// returns the first item of the deque without removing it
func (d *Deque) peekFront() int {
	if len(d.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}
	return d.items[0]
}

// returns the last item of the deque without removing it
func (d *Deque) peekBack() int {
	if len(d.items) == 0 {
		fmt.Println("Deque is empty")
		return -1
	}
	return d.items[len(d.items)-1]
}

// checks if the deque is empty
func (d *Deque) isEmpty() bool {
	return len(d.items) == 0
}

// returns the size of the deque
func (d *Deque) size() int {
	return len(d.items)
}

func main() {
	deque := Deque{}

	deque.addFront(5)
	deque.addFront(10)
	deque.addFront(15)
	deque.addFront(20)
	deque.addFront(25)

	fmt.Println(deque.removeFront())
	fmt.Println(deque)
	fmt.Println(deque.peekFront())
	fmt.Println(deque.peekBack())
	fmt.Println(deque.isEmpty())
	fmt.Println(deque.size())
}
