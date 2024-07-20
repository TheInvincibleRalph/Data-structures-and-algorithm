package main

import "fmt"

func heapSort(arr []int) {
	n := len(arr)
	for i = n/2 - 1; i >= 0; i-- { //creates a sub-tree: the loop begins by choosing a root index (i) then calls the heapify function to determine its left and right children and then produce a max heap eventually. The i-- ensures that we move from the last non-leaf node towards the root
		heapify(arr, n, i) //the heapify function takes the array, the length of the array, and the root index as parameters
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] //once a max heap has been formed by the heapify function, this moves the root value at index 0 to the end of the array and another heapification process continues afterward

		heapify(arr, i, 0) //this is a recursive function where index i keeps reducing by 1, thereby reducing the length of array ensuring that the largest numbers remain sorted at the rightmost end of each heapified array
	}

}

func heapify(arr []int, n, i int) {
	largest := i //initializes the largest value to be at the root index
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left //changes or updates the index number at largest to be the index at the left child
	}

	if right < n && arr[right] > arr[largest] {
		largest = right //if perhaps the value at the right-th index is greater than the value at the updated largest-th index, this updates it.
	}

	if largest != i { //this condition checks if largest index up to this stage is not the same as i (which is  n/2 - 1), if yes, the two indexes swap values
		arr[i], arr[largest] = arr[largest], arr[i]

		heapify(arr, n, largest) //this function continues until a max heap is produced.
	}

}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	heapSort(arr)

	fmt.Println("Sorted array is:", arr)
}
