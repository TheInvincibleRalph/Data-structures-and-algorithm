//This is a sliding window problem. We have to find the maximum sum of k consecutive elements in the array by first calculating the sum of the first k elements and then slide the window by one element to calculate the sum of the next k elements. Along the way, we will keep track of the maximum sum and return it as the answer.

package main

func maxSum(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1
	}

	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	maxSum := sum // initialize maxSum with the sum of the first k elements

	for i := k; i < n; i++ {
		sum += arr[i] - arr[i-k]
		//	arr[i] is the new element that is added to the window as it slides (in the first case being the kth element).
		// arr[i-k] is the element that is removed from the window as it slides.
		maxSum = max(maxSum, sum) // update maxSum if the current sum is greater
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 3
	println(maxSum(arr, k)) // Output: 13
}
