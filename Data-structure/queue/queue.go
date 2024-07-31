package main

import "fmt"

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

// returns the size of the queue
func (q *Queue) size() int {
	return len(q.items)
}

// returns the first item of the queue without removing it
func (q *Queue) peek() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func main() {
	queue := Queue{}

	queue.enqueue(5)
	queue.enqueue(10)
	queue.enqueue(15)
	queue.enqueue(20)
	queue.enqueue(25)

	fmt.Println(queue.dequeue())
	fmt.Println(queue)
	fmt.Println(queue.peek())
	fmt.Println(queue.isEmpty())
	fmt.Println(queue.size())
}
