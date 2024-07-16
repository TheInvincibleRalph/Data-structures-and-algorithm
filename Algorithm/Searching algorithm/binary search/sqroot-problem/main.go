package main

import "fmt"

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

func main() {
	fmt.Println(sqRoot(8))          // Output: 2
	fmt.Println(sqRoot(4))          // Output: 2
	fmt.Println(sqRoot(16))         // Output: 4
	fmt.Println(sqRoot(0))          // Output: 0
	fmt.Println(sqRoot(1))          // Output: 1
	fmt.Println(sqRoot(2147483647)) // Output: 46340 (edge case for large number)
}
