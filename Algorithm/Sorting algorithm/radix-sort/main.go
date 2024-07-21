package main

import "fmt"

func radixSort(arr []int) {
	max := maxDigit(arr) //this searches for the maximum number. This helps determine the number of digits to work with

	for exp := 1; max/exp > 0; exp *= 10 { //this performs counting sort for every digit, and this happens as long as max/exp is greater than zero. exp means exponent which refers to the place value of the digit in question and it increases by 10
		countSort(arr, exp)
	}
}

func maxDigit(arr []int) int {
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func countSort(arr []int, exp int) {

	n := len(arr)
	output := make([]int, n) //initializes output array
	count := make([]int, 10) //initializes count array to keep count of digit (from 0-9) occurences

	//count array
	for i := 0; i < n; i++ {
		count[(arr[i]/exp)%10]++ //this function extracts the digit at the current place value of arr[i] and increments that position in the initialized count array
	}

	//cumulative count array
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	//output array
	for i := n - 1; i >= 0; i-- { //starts iteration from the last element in the array (trasversing)
		output[count[(arr[i]/exp)%10]-1] = arr[i] //updates the output array with the value at each arr index
		count[(arr[i]/exp)%10]--                  //decrements cumulative count by 1 at every placement of arr value into the output array
	}

	// Copy the output array to arr[], so that arr now contains sorted numbers according to current digit
	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)
	radixSort(arr)
	fmt.Println("Sorted array:", arr)
}
