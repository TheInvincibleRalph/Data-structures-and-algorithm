package main

import "fmt"

//struct definition for other nodes
type Node struct {
	data int
	next *Node
}

//struct definition for the root node
type LinkedList struct {
	head *Node //the root node defined as "head" contains a data and a pointer to the next node
}

//adds element at the start of a linked list
func (list *LinkedList) addElementAtStart(data int) {
	newNode := &Node{data: data} //returns the memory addressnof Node struct so that the newNode variable can modify the struct directly
	newNode.next = list.head     //instantiate the original root to be the next in the list
	list.head = newNode          //replaces root node with the new node
}

//adds element at the end of a linkedlist
func (list *LinkedList) addElementAtEnd(data int) {
	newNode := &Node{data: data}
	current := list.head

	if list.head == nil {
		list.head = newNode
		return
	}
	for current.next != nil {
		current = current.next //the node originally pointed to by current is passed now to newNode
	}
	current.next = newNode //current.next points to newNode
}

// PrintList prints all elements in the list
func (list *LinkedList) PrintList() {
	current := list.head //instantiate the current node to be the head
	for current != nil {
		fmt.Printf("%d -> ", current.data) //loops through the list and print the data
		current = current.next             //update the current at every loop pass
	}
	fmt.Println("nil") //default value for the end of a linked list
}

func main() {
	list := LinkedList{}

	list.addElementAtStart(10)
	list.addElementAtStart(0)
	list.addElementAtEnd(20)
	list.addElementAtEnd(30)
	list.addElementAtEnd(40)
	list.addElementAtEnd(50)

	fmt.Println("Linked List:")
	list.PrintList()

}
