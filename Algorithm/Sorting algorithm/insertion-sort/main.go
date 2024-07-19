package main

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key { //for every j greater than or equal to zero and the value of j is greater than the current value of i
			arr[j+1] = arr[j] //Shift arr[j] to the right: updates the (j+1)th value to be the j-th value
			j--               //this will ensure the (j+1)th value (the one in the hole) above keeps moving left until no number is greater than it is.
		}
		arr[j+1] = key // since key was replaced with the j-th value this inserts it back into the j-th position (check concept.md for simulation)
		//once the nested for loop above breaks out, the key moves to the next index and the operation is performed all over again
	}

}
