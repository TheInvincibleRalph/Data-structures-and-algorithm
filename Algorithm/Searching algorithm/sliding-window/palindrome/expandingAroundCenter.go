package main

import "fmt"

func longestPalindrome(s string) string {
	start, maxLen := 0, 1

	//expandAroundCenter is a closure
	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen { //if the current window is greater than the maxLen
				start = left              //records the starting index of the current longest palindrome
				maxLen = right - left + 1 //updates the maxLen
			}
			left--  //expands the window by moving the left pointer a step backward
			right++ //expands the window by moving the right pointer a step forward
		}

	}
	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)   //odd length palindromes
		expandAroundCenter(i, i+1) //even length palindromes
	}

	return s[start : start+maxLen]
}

func main() {
	s := "forgeeksskeegfor"
	fmt.Println(longestPalindrome(s))
}
