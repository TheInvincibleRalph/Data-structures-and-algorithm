package main

type Items struct {
	value    int
	priority int
}

type PriorityQueue struct {
	items []Items
}

//<=========================================Helper Functions=========================================>

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// Swap two elements in the heap
func (pq *PriorityQueue) swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// <=================================Heapify Functions=================================>

func (pq *PriorityQueue) insert(value, priority int) {
	pq.items = append(pq.items, Items{value, priority})
	pq.heapifyUp(len(pq.items) - 1)
}

// Heapify up to maintain heap property
func (pq *PriorityQueue) heapifyUp(index int) {
	for index > 0 && pq.items[parent(index)].priority > pq.items[index].priority {
		pq.swap(index, parent(index))
		index = parent(index)
	}
}

// Extract the minimum element from the priority queue
func (pq *PriorityQueue) extractMin() (int, int) {
	if len(pq.items) == 0 {
		return -1, -1 // or some error value
	}
	min := pq.items[0]
	pq.items[0] = pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.heapifyDown(0)
	return min.value, min.priority
}

func (pq *PriorityQueue) heapifyDown(i int) {
	for {
		left := left(i)
		right := right(i)
		smallest := i //initialize smallest to the parent

		//check if the left child is less than the parent
		if left < len(pq.items) && pq.items[left].priority < pq.items[smallest].priority {
			smallest = left
		}

		//check if the right child is less than the parent
		if right < len(pq.items) && pq.items[right].priority < pq.items[smallest].priority {
			smallest = right
		}

		//if the parent is less than the children, then the heap is valid
		if i == smallest {
			break
		}

		//swap the parent with the smallest child
		pq.swap(i, smallest)
		i = smallest

	}
}

// Peek the minimum element without removing it
func (pq *PriorityQueue) peek() (int, int) {
	if len(pq.items) == 0 {
		return -1, -1 // or some error value
	}
	return pq.items[0].value, pq.items[0].priority
}
