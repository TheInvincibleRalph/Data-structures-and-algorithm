package main

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

//when a circular queue is instantiated, it returns a pointer to a CircularQueue struct with with an added slot to the capacity the queue was instantiated with
