package Array

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
