package main

import "fmt"

func minMax(arr []int) (int, int) {
	min, max := arr[0], arr[0] //initialization min and max value

	for _, value := range arr { //ranges over the array "arr" and stores the result in "value"
		if value < min { //at each iteration, if value is less than min, update min, same goes for max when value is greater than max
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {
	arr := []int{14, 10, 55, 83, 63, 5, 17}
	min, max := minMax(arr)
	fmt.Println(min)
	fmt.Println(max)
}

/*
When iterating over an array with range,
two values are returned for each iteration:
the index and the value of the element at that index.
The underscore (_) is used to ignore the index since it's not needed in this particular operation.
*/
