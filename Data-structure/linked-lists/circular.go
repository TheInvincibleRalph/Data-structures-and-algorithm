package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type CircularLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (cll *CircularLinkedList) display() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}
	temp := cll.head
	for {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
		if temp == cll.head {
			break
		}
	}
	fmt.Println("Head")
}

func (cll *CircularLinkedList) addFromStart(data int) {
	newNode := &Node{data: data}

	if cll.head == nil {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
		newNode.previous = newNode
	}
	newNode.next = cll.head
	cll.tail.next = newNode
	newNode.previous = cll.tail
	cll.head.previous = newNode
	cll.head = newNode
}

func (cll *CircularLinkedList) addFromEnd(data int) {
	newNode := &Node{data: data}

	if cll.head == nil {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
		newNode.previous = newNode
		return
	}
	newNode.next = cll.head
	newNode.previous = cll.tail
	cll.tail.next = newNode
	cll.head.previous = newNode
}

func (cll *CircularLinkedList) addFromWithin(nodeData int, newData int) {
	newNode := &Node{data: newData}

	if cll.head == nil {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
		newNode.previous = newNode
	}

	current := cll.head
	for current != nil && current.data != nodeData {
		current = current.next
	}
	newNode.next = current.next
	newNode.previous = current
	current.next.previous = newNode
	current.next = newNode
}
func main() {
	list := CircularLinkedList{}

	list.addFromStart(40)
	list.addFromStart(30)
	list.addFromStart(20)
	list.addFromStart(10)
	list.addFromStart(0)
	list.addFromEnd(50)
	list.addFromWithin(10, 15)
	fmt.Println("List is:")
	list.display()
}
