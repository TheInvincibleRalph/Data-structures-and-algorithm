package main

type Heap struct {
	items []int
}

// <=================================Helper Functions=================================>

//returns the index of the parent node
func parent(i int) int {
	return (i - 1) / 2
}

//returns the index of the left child
func left(i int) int {
	return 2*i + 1
}

//returns the index of the right child
func right(i int) int {
	return 2*i + 2
}

//swaps two items within a slice
func (h *Heap) swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}
