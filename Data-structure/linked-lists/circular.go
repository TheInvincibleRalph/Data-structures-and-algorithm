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

func (cll *CircularLinkedList) traverse() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := cll.tail
	for {
		fmt.Printf("%d -> ", current.data)
		current = current.previous
		if current == cll.head.previous {
			break
		}
	}
	fmt.Println("Tail")
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
	cll.tail = newNode //updates the tail pointer

}

func (cll *CircularLinkedList) addFromWithin(nodeData int, newData int) {
	newNode := &Node{data: newData}

	if cll.head == nil {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
		newNode.previous = newNode
		return
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

func (cll *CircularLinkedList) deleteFromStart() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if cll.head.next == cll.head {
		cll.head = nil // Only one node
		cll.tail = nil
		return
	}

	cll.head.next.previous = cll.tail
	cll.tail.next = cll.head.next
	cll.head = cll.head.next

}

func (cll *CircularLinkedList) deleteFromEnd() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if cll.head.next == cll.head {
		cll.head = nil // Only one node
		cll.tail = nil
		return
	}
	cll.head.previous = cll.tail.previous
	cll.tail.previous.next = cll.head
	cll.tail = cll.tail.previous
}

func (cll *CircularLinkedList) delByValue(value int) {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if cll.head.next == cll.head {
		cll.head = nil // Only one node
		cll.tail = nil
		return
	}

	current := cll.head
	for current != nil && current.next.data != value {
		current = current.next
		if current == cll.tail {
			break
		}
	}
	current.next = current.next.next
	current.next.previous = current
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
	list.deleteFromStart()
	list.deleteFromEnd()
	list.delByValue(15)
	list.traverse()
	fmt.Println("List is:")
	list.display()

}
