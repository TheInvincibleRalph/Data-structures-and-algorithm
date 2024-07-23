package main

import "fmt"

func deletion(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Index out of bound")
		return arr
	}

	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	arr = deletion(arr, 2)
	fmt.Println("Array after deletion:", arr)
}
