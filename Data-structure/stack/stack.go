package main

type Stack struct {
	items []int
}

// this adds an item to the stack
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}
