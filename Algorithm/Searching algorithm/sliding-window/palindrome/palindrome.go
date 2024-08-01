//A palindrome is a string that reads the same forward and backward.
package main

import "fmt"

func palindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "racecar"
	fmt.Println(palindrome(s))
}
