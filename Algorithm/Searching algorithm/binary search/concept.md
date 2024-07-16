## What is Binary Search?

In computer science, binary search, also known as half-interval search, logarithmic search, or binary chop, is a search algorithm that finds the position of a target value within a sorted array. Binary search compares the target value to the middle element of the array. If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half, again taking the middle element to compare to the target value, and repeating this until the target value is found. If the search ends with the remaining half being empty, the target is not in the array. [Wikipedia](https://en.wikipedia.org/wiki/Binary_search)



## How does it work?

In its simplest form, Binary Search operates on a contiguous sequence with a specified left and right index. This is called the Search Space. Binary Search maintains the left, right, and middle indicies of the search space and compares the search target or applies the search condition to the middle value of the collection; if the condition is unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful. If the search ends with an empty half, the condition cannot be fulfilled and target is not found. [Leetcode](https://leetcode.com/explore/learn/card/binary-search/138/background/971/)


## Sqaure-root Problem

```go
package main

import "fmt"

// mySqrt function takes a non-negative integer x and returns the floor value of its square root.
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }

    left, right := 0, x
    var result int

    // Perform binary search
    for left <= right {
        mid := left + (right - left) / 2
        
        // Compare mid*mid with x
        if mid * mid == x {
            return mid // Exact square root found
        } else if mid * mid < x {
            left = mid + 1
            result = mid // Update result to the last valid mid
        } else {
            right = mid - 1
        }
    }

    return result // Return the floor value of the square root
}

func main() {
    // Test cases
    fmt.Println(mySqrt(4))  // Output: 2
    fmt.Println(mySqrt(8))  // Output: 2
    fmt.Println(mySqrt(16)) // Output: 4
    fmt.Println(mySqrt(0))  // Output: 0
    fmt.Println(mySqrt(1))  // Output: 1
    fmt.Println(mySqrt(2147483647)) // Output: 46340 (edge case for large number)
}
```

### Conceptual example

Consider finding the square root of \( x = 10 \):

1. **Initial Range**:
   - `left = 0`
   - `right = 10`

2. **First Iteration**:
   - `mid = (0 + 10) / 2 = 5`
   - `5 * 5 = 25` which is greater than 10, so adjust the range to `[0, 4]`.

3. **Second Iteration**:
   - `mid = (0 + 4) / 2 = 2`
   - `2 * 2 = 4` which is less than 10, so adjust the range to `[3, 4]`.
   - Update `result = 2`.

4. **Third Iteration**:
   - `mid = (3 + 4) / 2 = 3`
   - `3 * 3 = 9` which is less than 10, so adjust the range to `[4, 4]`.
   - Update `result = 3`.

5. **Fourth Iteration**:
   - `mid = (4 + 4) / 2 = 4`
   - `4 * 4 = 16` which is greater than 10, so adjust the range to `[4, 3]`.

6. **Termination**:
   - The search range is invalid (`left > right`), so the loop terminates and returns `result = 3`.


#### References

In the sqRoot function: That particular formula is used because when using (left + right) / 2, if left and right are large, their sum left + right could exceed the maximum value that can be stored in an integer variable, leading to an overflow. On the other hand, left + (right - left) / 2 ensures that no overflow occurs because right - left is calculated first, which is always a non-negative value and much smaller than left + right.