package main

type Queue struct {
	items []int
}

//enqueue adds an item to the queue
func (q *Queue) enqueue(item int) {
	q.items = append(q.items, item)
}

//dequeue removes the first item added
func (q *Queue) dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
