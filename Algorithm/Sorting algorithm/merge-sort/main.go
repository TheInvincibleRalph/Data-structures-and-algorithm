package main

import "fmt"

func mergeSort(arr []int) {

	if len(arr) > 1 {

		mid := len(arr) / 2

		// Create copies of the slices to avoid modifying the same underlying array
		left := append([]int(nil), arr[:mid]...)
		right := append([]int(nil), arr[mid:]...)

		mergeSort(left)  //recursive call of all left arrays
		mergeSort(right) //recursive call of all right arrays

		i, j, k := 0, 0, 0

		for i < len(left) && j < len(right) { //compare elements from left and right subarrays
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j] //since it will be the smallest if left[i] is greater
				j++
			}
			k++
		}

		// Copy any remaining elements from the left subarray
		for i < len(left) {
			arr[k] = left[i]
			i++
			k++
		}

		// Copy any remaining elements from the right subarray
		for j < len(right) {
			arr[k] = right[j]
			j++
			k++
		}
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	mergeSort(arr)
	fmt.Println("Sorted array is:", arr)
}
