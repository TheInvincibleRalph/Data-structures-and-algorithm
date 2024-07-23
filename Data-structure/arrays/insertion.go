package main

import "fmt"

func insertion(arr []int, index int, value int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Index out of bound")
		return arr
	}

	sliceBeforeIndex := arr[:index] //this creates a slice of array that contains elements before the index (the element at the index exclusive)
	sliceAfterIndex := arr[index:]  //this creates a slice of array that contains elements after the array (the element at the index inclusive)

	arr = append(sliceBeforeIndex, append([]int{value}, sliceAfterIndex...)...) //an array is first created internally with the slice of elements after the index appended to the element at the given index, the result is afterwards appended to the slice before the index
	return arr
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	arr = insertion(arr, 2, 25) // Insert 25 at index 2
	fmt.Println("Array after insertion:", arr)
}
