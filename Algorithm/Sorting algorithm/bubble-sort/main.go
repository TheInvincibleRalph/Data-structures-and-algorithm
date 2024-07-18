package main

import "fmt"

func bubbleSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap if elements are out of order
			}
		}
	}

}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	bubbleSort(arr)  // Call the function to sort the array
	fmt.Println(arr) // The array is now sorted
}
