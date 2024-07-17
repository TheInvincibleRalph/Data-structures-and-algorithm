package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1 //handle edge case of empty array
	}

	maxElement := arr[0] //initialize the maximum element to be the first element of the array.
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxElement {
			maxElement = arr[i] //updating maximum element to be the current value
		}
	}
	return maxElement
}

func main() {
	arr := []int{1, 4, 45, 6, 0, 19}
	maxNum := findMax(arr)
	fmt.Println("Maximum element in the array is", maxNum) // Output: 45
}
