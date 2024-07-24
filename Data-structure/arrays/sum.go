package main

import "fmt"

func sumOf(arr []int) int {
	sum := 0

	for _, element := range arr {
		sum = sum + element //or sum += element
	}
	return sum
}

func main() {
	arr := []int{14, 10, 55, 83, 63, 5, 17}
	sum := sumOf(arr)
	fmt.Println(sum)
}
