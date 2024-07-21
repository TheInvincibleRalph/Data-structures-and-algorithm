package main

import "fmt"

func countSort(arr []int) []int {
	//get maximum value in the array
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// fmt.Println(max)

	//create and initialize the count array
	n := len(arr)
	count := make([]int, max+1)
	// fmt.Println(output)
	// fmt.Println(count)

	//counting array
	for _, num := range arr {
		count[num]++
		// fmt.Println(count)
	}

	//cumulative count
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
		// fmt.Println(count)
	}

	//build output array
	output := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
		// fmt.Println(count)
		// fmt.Println(output)
	}

	// Copy the output array to the original array
	copy(arr, output)
	return arr
}

func main() {
	arr := []int{4, 2, 2, 8, 3, 3, 1}
	sortedArray := countSort(arr)
	fmt.Println("Sorted array is:", sortedArray)
}

//NB: Feel free to uncomment the commented print functions to see the result of each process under the hood.
