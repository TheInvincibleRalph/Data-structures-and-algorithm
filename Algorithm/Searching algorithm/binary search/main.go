/*
Question

Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index.
Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity. //this means the maximum number of operations needed to search the target number is a logarithmic function of the size of the search space.

*/

package main

import "fmt"

func binSearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		m := (left + right) / 2
		if arr[m] < num {
			left = m + 1
		} else if arr[m] > num {
			right = m - 1
		} else {
			return m
		}
	}
	return -1 // Element not found
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	num := 10
	result := binSearch(arr, num)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}
