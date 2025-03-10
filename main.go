package main

import (
	"dsa/Queue"
	"fmt"
)

func main() {

	// String :

	// String.CheckPalindrome("malayalam")
	// String.FindfirstNonRepeating("swiss")
	// String.ReverseString("Hello")
	// String.CountVowels("Hello World")

	// Array :
	// array := []int{1, 2, 3, 4, 5, 2, 8, 5, 3}
	// Array.ReverseArray(array)
	// Array.FindLargestElement(array)
	// Array.FindDuplicates(array)
	// Array.RemoveDuplicates(array)

	// LinkedList :
	// list := LinkedList.LinkedList{}
	// list.InsertBeginning(23)
	// list.InsertAtEnd(80)
	// list.DeleteFromBeginning()
	// list.DeleteFromEnd()
	// list.FindElementLinkedList(44)
	// list.FindMiddleElement()
	// list.LegthofLinkedList()
	// list.ReverseLinkedList()
	// list.Display()

	// Stack :
	// myStack := Stack.Stack{}
	// myStack.Push(80)
	// myStack.Push(70)
	// res := myStack.Pop()
	// fmt.Println("Removed item:", res)
	// fmt.Println(myStack)

	// Queue :
	myQueue := Queue.Queue{}
	myQueue.Enqueue(80)
	myQueue.Enqueue(70)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
