package LinkedList

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

// Insert Node at Beginning
func (list *LinkedList) InsertBeginning(value int) {
	newNode := &Node{Data: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.Next = list.head
	list.head = newNode
}

// Insert Node at End
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

// Delete a Node at Beginning
func (list *LinkedList) DeleteFromBeginning() {
	if list.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	if list.head.Next == nil {
		list.head = nil
		return
	}
	list.head = list.head.Next

}

// Delete a Node at End
func (list *LinkedList) DeleteFromEnd() {
	if list.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	if list.head.Next == nil {
		list.head = nil
		return
	}
	current := list.head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

// Length of LinkedList
func (list *LinkedList) LegthofLinkedList() {
	length := 0
	current := list.head
	for current != nil {
		length++
		current = current.Next
	}
	fmt.Println("Length of linkedlist :", length)
}

// Find an Element
func (list *LinkedList) FindElementLinkedList(element int) {
	current := list.head
	length := 0
	for current != nil {
		length++
		if current.Data == element {
			fmt.Printf("Element Found on %d th node\n", length)
			return
		}
		current = current.Next
	}
	fmt.Println("Element not found !")
}

// Reverse a LinkedList
func (list *LinkedList) ReverseLinkedList() {
	var prev *Node
	current := list.head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.head = prev
}

// Find MiddleElement of LinkedList
func (list *LinkedList) FindMiddleElement() {
	slow := list.head
	fast := list.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Println("Middle Element :", slow.Data)
}
