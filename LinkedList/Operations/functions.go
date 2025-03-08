package operations

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertBeginning(value int) {
	newNode := &Node{Data: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.Next = list.head
	list.head = newNode
}
func (list *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{Data: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (list *LinkedList) DeleteFromBeginning() {
	if list.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	list.head = list.head.Next
	return
}
