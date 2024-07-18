package main

import "fmt"

var pick int

//This is a mock API that once called determines the nature of the guessed number
//It takes in the mid value from the binary search function (guessNumber)
func guessAPI(mid int) int {

	if mid > pick {
		return -1
	} else if mid < pick {
		return 1
	} else {
		return 0
	}
}

//The user makes continous guess within a search space bounded by n and the guessAPI uses the mid-value to determine if the guess is greater, lesser, or equal to the picked number.
//NB: The for loop operation simulates the user's continuous guess

func guessNumber(n int) int {
	left, right := 1, n //here, n defines the upper bound of the search range and the same as the argument to the function

	for left <= right {
		mid := left + (right-left)/2

		if guessAPI(mid) == -1 {
			right = mid - 1
		} else if guessAPI(mid) == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	pick = 6
	n := 10

	fmt.Println("The picked number is", guessNumber(n))
}
