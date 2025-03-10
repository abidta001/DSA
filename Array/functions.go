package Array

/* Array is a collection of elements into a single variable of same type.
Basically array have a definite size or we can use slice or variadic parameter[...] for making array dynamic */

import "fmt"

// Reverse Array
func ReverseArray(array []int) []int {
	i, j := 0, len(array)-1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i, j = i+1, j-1
	}
	return array
}

// Find Largest Element
func FindLargestElement(array []int) int {
	largest := array[0]
	for i := range array {
		if array[i] > largest {
			largest = array[i]
		}
	}
	return largest
}

// Find Duplicates
func FindDuplicates(array []int) {
	slice := []int{}
	newMap := make(map[int]int)
	for _, i := range array {
		newMap[i]++
	}
	for i, j := range newMap {
		if j > 1 {
			slice = append(slice, i)
		}
	}
	fmt.Println("Duplicates: ", slice)
}

// Remove Duplicates(with Duplicate occuring once)
func RemoveDuplicates(array []int) {
	slice := []int{}
	newMap := make(map[int]int)
	for _, i := range array {
		newMap[i]++
	}
	for i, _ := range newMap {
		slice = append(slice, i)
	}
	fmt.Println("Duplicates removed : ", slice)
}
