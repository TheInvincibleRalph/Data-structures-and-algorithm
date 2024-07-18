package main

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1                   //hole
		for j >= 0 && arr[j] > key { //for every j greater than or equal to zero and the value of j is greater than the current value of i
			arr[j+1] = arr[j] //move the value in j to the right by moving the (j+1)th value to the left
			j--               //repeat the same recursive swap as long as the condition permits
		}
		arr[j+1] = key //once the condition is exhausted, that is, the j-th value in the array is now less than the current i-th value, the least of the values is left in the (j+1)th position, it then replaces the current key
	}
}
