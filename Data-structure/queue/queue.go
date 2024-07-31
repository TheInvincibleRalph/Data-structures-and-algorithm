package main

type Queue struct {
	items []int
}

// enqueue adds an item to the queue
func (q *Queue) enqueue(item int) {
	q.items = append(q.items, item)
}

// dequeue removes the first item added
func (q *Queue) dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// isEmpty checks if the queue is empty
func (q *Queue) isEmpty() bool {
	return len(q.items) == 0 //returns true if the length of the queue is 0. This syntax is idiomatic for booleans in Go
}
