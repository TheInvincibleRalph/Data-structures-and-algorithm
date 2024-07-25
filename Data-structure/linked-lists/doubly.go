package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (dll *DoublyLinkedList) addFromStart(data int) {
	newNode := &Node{data: data}
	current := dll.head

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = current
		current.previous = newNode
		dll.head = newNode
	}
	dll.length++

}

func (dll *DoublyLinkedList) printList() {
	if dll.head == nil {
		fmt.Println("Empty list")
		return
	}

	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
	fmt.Println("Length of list is ", dll.length)
}

func main() {
	list := DoublyLinkedList{}

	list.addFromStart(10)
	list.addFromStart(0)
	list.addFromStart(20)
	fmt.Println("Linked List:")
	list.printList()

}
