/*
Question
Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index.
Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity. //this means the maximum number of operations needed to search the target number is a logarithmic function of the size of the search space.

Constraints
1) 1 <= len(arr) <= 104
2) -104 < arr[i], num < 104
3) All the integers in nums are unique.
4) nums is sorted in ascending order.
*/

package main

import "fmt"

//this function takes a sorted array nums and an integer target as inputs and returns the index of the target if found, otherwise returns -1.

func binSearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1

	for left <= right { //this loop continues as long as there is a valid search space (a constraint imposed by left <= right)
		//recalculates m each iteration
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
