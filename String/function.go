package String

import "fmt"

// Check Palindrome
func CheckPalindrome(s string) {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			fmt.Printf("%s is not palindrome\n", s)
		}
		i, j = i+1, j-1
	}
	fmt.Printf("%s is palindrome\n", s)
}

// First Non-repeating character :
func FindfirstNonRepeating(s string) {
	newMap := make(map[rune]int)
	for _, i := range s {
		newMap[i]++
	}
	for _, i := range s {
		if newMap[i] == 1 {
			fmt.Println("First Non repeating element is :", string(i))
			return
		}
	}
}
