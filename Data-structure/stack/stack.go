package main

import "fmt"

type Stack struct {
	items []int
}

// this adds an item to the stack
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	stack := Stack{}

	stack.push(5)
	stack.push(10)
	stack.push(15)
	stack.push(20)
	stack.push(25)
	stack.pop()

	fmt.Println(stack)
}
