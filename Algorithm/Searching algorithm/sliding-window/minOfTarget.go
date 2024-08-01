package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Frequency maps
	target := make(map[byte]int) //stores the frequency of characters in t
	window := make(map[byte]int) //stores the frequency of characters in the window

	for i := range t {
		target[t[i]]++ //populate the target map with the frequency of characters in t
	}

	have, need := 0, len(target) //have is the number of characters in target that are in window, need is the number of characters in target
	left, right := 0, 0          //left and right pointers
	minLength := len(s) + 1
	start := 0 //records the start of the substring

	// checks if character in the window is in target
	for right < len(s) { //(expansion process)
		char := s[right]
		window[char]++

		if window[char] == target[char] {
			have++ // increment have if the character is in target
		}

		for have == need { // if all characters in target are in window
			if right-left+1 < minLength { // if the window is smaller than the current minimum
				minLength = right - left + 1
				start = left //records the starting index of the current smallest window that contins all characters from t
			}

			// If the frequency in the window is
			// now less than required by target,
			// it means that the window no longer
			// meets the requirement for this character.
			// Therefore, the have counter is decremented
			// to reflect that not all required characters
			// are present in the window anymore.
			// The left pointer is then moved to the right
			window[s[left]]--                      //decrement the frequency of the character at the left pointer from the window
			if window[s[left]] < target[s[left]] { //if frquency of a particular character in windows is less than its frequency in target, decrement
				have--
			}
			left++ //(shrink process)
		}

		right++
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[start : start+minLength] // return the substring
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}

//NB: When you seed a map with a key, the output will be the 'value' that is why line 29 makes sense.
