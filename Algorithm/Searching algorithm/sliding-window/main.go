package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(arr []int, x int) int {
	if len(arr) == 0 {
		return -1 //handle edge case of empty array
	}

	start := 0                 //instantiate the start index
	minLength := math.MaxInt32 //instantiate a minimum length variable and set it to "infinity"
	currentSum := 0            //set current sum to zero

	//this loop function expands the window
	for end := 0; end < len(arr); end++ {
		currentSum += arr[end]

		//this loop function checks for sum condition and shrinks the window if condition is met
		for currentSum > x {
			minLength = min(minLength, end-start+1)
			currentSum -= arr[start]
			start++
		}
	}

	if minLength == math.MaxInt32 {
		return 0 //no subarray found
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 4, 45, 6, 0, 19}
	x := 51
	fmt.Println(minSubArrayLen(arr, x))
}
