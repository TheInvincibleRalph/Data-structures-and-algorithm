package main

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // map to store the index of the last occurrence of a character
	start, maxLength := 0, 0        // start is the starting index of the substring, maxLength is the length of the longest substring

	for end := 0; end < len(s); end++ { // end is the ending index of the substring
		if index, found := charIndex[s[end]]; found && index >= start { // if the character is found in the map and the index is greater than or equal to the start index
			start = index + 1 // update the start index to the index of the character + 1
		}
		charIndex[s[end]] = end                 // update the index of the character in the map
		maxLength = max(maxLength, end-start+1) // update the maxLength
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
