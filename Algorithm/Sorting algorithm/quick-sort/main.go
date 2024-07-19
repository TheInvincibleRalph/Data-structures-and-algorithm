package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSorter(arr, 0, len(arr)-1)

}

// quickSorter is the recursive function that sorts the subarrays.
func quickSorter(arr []int, low, high int) {

	if low < high {
		p := partition(arr, low, high)
		quickSorter(arr, low, p-1)  //the array at left-hand side of the pivot element with its low and high indexes
		quickSorter(arr, p+1, high) //the array at the right-hand side of the pivot element with its low and high indexes
	}
}

// partition function rearranges the elements based on the pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++ //continues to swap the i-th and j-th elements as long as the condition permits
		}
	}
	//the for loop ends in a state where all element to the left of i is less than the pivot element such that the pivot swapping position with the i-th element makes sense
	arr[i], arr[high] = arr[high], arr[i]
	return i // the next pivot point is returned and the sorting process starts all over.
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	quickSort(arr)
	fmt.Println("Sorted array is:", arr)
}
