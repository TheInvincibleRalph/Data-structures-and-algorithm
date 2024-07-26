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
		dll.head = newNode //updates the head node
	}
	dll.length++
}

func (dll *DoublyLinkedList) addFromEnd(data int) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.previous = dll.tail
		dll.tail = newNode //updates the tail node
	}
	dll.length++
}

func (dll *DoublyLinkedList) addFromWithin(newData int, nodeData int) {
	newNode := &Node{data: newData}
	current := dll.head

	if dll.head == nil {
		dll.head = newNode
		return
	}

	for current != nil && current.data != nodeData {
		current = current.next
	}
	newNode.next = current.next
	current.next.previous = newNode
	current.next = newNode
	newNode.previous = current
	dll.length++
}

func (dll *DoublyLinkedList) delFromStart() {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}
	dll.head = dll.head.next
	if dll.head != nil {
		dll.head.previous = nil
	} else {
		dll.tail = nil
	}
	dll.length--
}

func (dll *DoublyLinkedList) delFromEnd() {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if dll.head.next == nil {
		dll.head = nil
		dll.tail = nil
		return
	}

	current := dll.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	dll.tail = current
	dll.length--

}

func (dll *DoublyLinkedList) delByValue(value int) {
	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := dll.head

	// Check if the node to delete is the head node
	if dll.head.data == value {
		dll.delFromStart()
		return
	}

	for current.next != nil && current.next.data != value {
		current = current.next
	}

	// If the node is not found, return
	if current.next == nil {
		fmt.Println("Value not found in list")
		return
	}

	current.next = current.next.next

	if current.next != nil {
		current.next.previous = current
	} else {
		dll.tail = current
	}
	dll.length--

}

func (dll *DoublyLinkedList) printForward() {
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

func (dll *DoublyLinkedList) printBackward() {
	if dll.tail == nil {
		fmt.Println("Empty list")
		return
	}

	current := dll.tail

	for current != nil {
		fmt.Printf("%d <- ", current.data)
		current = current.previous
	}
	fmt.Println("nil")
	fmt.Println("Length of list is ", dll.length)
}

func main() {
	list := DoublyLinkedList{}

	list.addFromStart(10)
	list.addFromStart(0)
	list.addFromStart(20)
	list.addFromEnd(40)
	list.addFromWithin(30, 10)
	list.delFromStart()
	list.delFromEnd()
	list.delByValue(10)
	fmt.Println("Linked List:")
	list.printForward()
	list.printBackward()

}

/*
Ensure that when nodes are added to a list,
all fields (data, previous, next) are properly
initialized for data integrity sake.
*/
