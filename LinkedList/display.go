package LinkedList

import "fmt"

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("Empty List ...")
		return
	}
	current := list.head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
}
