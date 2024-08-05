package main

import "fmt"

type CircularQueue struct {
	items    []int
	front    int
	rear     int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]int, capacity+1),
		front:    0,
		rear:     0,
		capacity: capacity + 1,
	}
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == cq.rear
}

func (cq *CircularQueue) IsFull() bool {
	return (cq.rear+1)%cq.capacity == cq.front
}

func (cq *CircularQueue) Enqueue(item int) bool {
	if cq.IsFull() {
		return false
	}
	cq.items[cq.rear] = item              //cq.rear is always pointing to the next empty slot
	cq.rear = (cq.rear + 1) % cq.capacity //cq.rear is updated to the next empty slot within the capacity of the queue
	return true
}

func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	item := cq.items[cq.front]              //cq.front is always pointing to the first element of the queue
	cq.front = (cq.front + 1) % cq.capacity //cq.front is updated to the next element in the queue
	return item, true
}

func (cq *CircularQueue) Front() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	return cq.items[cq.front], true
}

func (cq *CircularQueue) Rear() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	return cq.items[(cq.rear+cq.capacity-1)%cq.capacity], true
}

func main() {
	cq := NewCircularQueue(5) //when a circular queue is instantiated, it returns a pointer to a CircularQueue struct with with an added slot to the capacity the queue was instantiated with

	cq.Enqueue(1)
	cq.Enqueue(2)
	cq.Enqueue(3)
	cq.Enqueue(4)
	cq.Enqueue(5)
	cq.Enqueue(6) //this will return false as the queue is full
	cq.Dequeue()
	cq.Enqueue(6) //this will return true as the queue has an empty slot after the last element

	for !cq.IsEmpty() {
		fmt.Println(cq.Dequeue())
	}
}
