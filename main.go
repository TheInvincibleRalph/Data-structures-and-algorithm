/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.
*/

package main

import "fmt"

func guessNum(guess int) int {

	left, right := 1, 100
	num := 10

	for i := 0; i < 100; i++ {
		//mid = left + (right-left)/2

		if guess > num {
			return -1
		} else if guess < num {
			return left
		} else if guess == num {
			return num
		}
	}
	return right
}

func main() {
	fmt.Println(guessNum(10))
}
