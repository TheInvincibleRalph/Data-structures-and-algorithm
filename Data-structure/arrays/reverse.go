package main

import "fmt"

func reverseArray(arr [5]int) [5]int {
	n := len(arr)
	for i := 0; i < (n / 2); i++ {
		j := n - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	reversedArr := reverseArray(arr)
	fmt.Println("Reversed array:", reversedArr) // Output: [50, 40, 30, 20, 10]
}
