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

	fmt.Println(sqRoot(8))          // Output: 2
	fmt.Println(sqRoot(4))          // Output: 2
	fmt.Println(sqRoot(16))         // Output: 4
	fmt.Println(sqRoot(0))          // Output: 0
	fmt.Println(sqRoot(1))          // Output: 1
	fmt.Println(sqRoot(2147483647)) // Output: 46340 (edge case for large number)
}

func sqRoot(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x //variable declaration and assignment, and a search space of square root of any positive number
	var result int

	for left <= right {
		mid := left + (right-left)/2 //see concept.md for why this formula is used to calculate mid

		//Based on the comparison, the search range is adjusted by moving the
		//left or right pointers, effectively narrowing down the
		//potential values for the square root.
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1

		}
	}
	return result //returns the largest integer whose square is less than or equal to x
}
