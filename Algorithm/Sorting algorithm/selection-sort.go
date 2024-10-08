// Principle: Select the smallest (or largest) element from the unsorted part and swap it with the first unsorted element.

package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minVal := i //initializes the index of the minimum value to be the current index of i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minVal] { //condition to check if the value at index j is less than the initial value at index minIndex
				minVal = j //updates the minVal index to be the current index as approved by the if condition.
			}

		}
		arr[i], arr[minVal] = arr[minVal], arr[i] //\swap the elements or values in the indexes
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}

	selectionSort(arr)
	fmt.Println("The sorted array is:", arr)
}

/*
On using minVal

Using minVal (or a similar variable) correctly maintains the state
of the smallest value found so far during iterations.
Using i instead would lead to incorrect comparisons and results,
as i changes with each iteration and doesn't preserve the required
state for finding the minimum or maximum element efficiently.
In algorithms where tracking indices or values is critical,
using a dedicated variable ensures correct and consistent operations.
*/
