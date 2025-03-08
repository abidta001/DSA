package main

import (
	operations "dsa/LinkedList/Operations"
	"fmt"
)

func main() {
	list := &operations.LinkedList{}
	list.InsertBeginning(25)
	list.InsertBeginning(90)
	list.InsertBeginning(44)
	list.InsertAtEnd(90)
	list.Display()
	list.DeleteFromBeginning()
	fmt.Println()
	list.Display()
}
