# Nested Loop Analysis


## Case Study

```go
func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {      // Outer loop
        for j := 0; j < n-i-1; j++ { // Inner loop
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // Swap elements if they are in the wrong order
            }
        }
    }
}
```

### Explanation of `i` and `j`

#### Outer Loop (`i`)
- **Purpose**: The outer loop, controlled by `i`, represents the number of passes over the array. Each pass ensures that the largest element in the unsorted part of the array is moved to its correct position.
- **Range**: `i` runs from `0` to `n-2`. This means the loop runs `n-1` times.
- **Why `n-1`?**: After `n-1` passes, the array is sorted. In each pass, one element is placed in its correct position. By the time `i` reaches `n-2`, the second-last element is already in its correct position after `n-1` comparisons.

#### Inner Loop (`j`)
- **Purpose**: The inner loop, controlled by `j`, performs the actual comparisons and swaps of adjacent elements if they are in the wrong order.
- **Range**: `j` runs from `0` to `n-i-2`. This range ensures that the loop does not compare the already sorted elements. The inner loop runs `n-i-1` times.
- **Why `n-i-1`?**: After each pass of the outer loop, the last `i` elements-*element(s) at the rightmost position of the array*-are in their correct positions. Hence, the inner loop runs only up to `n-i-2`-which makes it `n-i-1` times-to avoid unnecessary comparisons with already sorted elements.

### How `i` and `j` Work Together

1. **First Pass (`i = 0`)**:
   - `j` runs from `0` to `n-2` *which is from range `0` to (`n-i-2).*
   - Compares and potentially swaps every adjacent pair from the start to the second-to-last element.
   - The largest element bubbles up to the end of the array.

2. **Second Pass (`i = 1`)**:
   - `j` runs from `0` to `n-3` *which is from range `0` to (`n-i-2).*
   - The second largest element bubbles up to its correct position just before the largest element.

3. **Subsequent Passes**:
   - As `i` increases, the range of `j` decreases.
   - This is because the last `i` elements are already sorted, and we don't need to compare them again.

### Concept simulation

Consider an array: `[5, 3, 8, 4, 2]`

1. **First Pass (`i = 0`)**:
   - `j` runs from `0` to `3` and the number decreases as long as j is less than `n-i-1`
   - Comparisons: `(5, 3)`, `(5, 8)`, `(8, 4)`, `(8, 2)`
   - Array after first pass: `[3, 5, 4, 2, 8]`

2. **Second Pass (`i = 1`)**:
   - `j` runs from `0` to `2` and the number decreases as long as j is less than `n-i-1`
   - Comparisons: `(3, 5)`, `(5, 4)`, `(5, 2)`
   - Array after second pass: `[3, 4, 2, 5, 8]`

3. **Third Pass (`i = 2`)**:
   - `j` runs from `0` to `1` and the number decreases as long as j is less than `n-i-1`
   - Comparisons: `(3, 4)`, `(4, 2)`
   - Array after third pass: `[3, 2, 4, 5, 8]`

4. **Fourth Pass (`i = 3`)**:
   - `j` runs from `0` to `0`.
   - Comparison: `(3, 2)`
   - Array after fourth pass: `[2, 3, 4, 5, 8]`

After four passes, the array is sorted, and the outer loop completes.

### Summary

- **`i`**: Tracks the number of passes and helps in reducing the range of `j` to avoid unnecessary comparisons with already sorted elements.
- **`j`**: Performs the actual comparisons and swaps to sort the array during each pass.
- Together, they ensure that the largest elements bubble up to their correct positions after each pass, progressively sorting the array.

**NB**: In a nested loop, each time we execute the outer loop, the inner loop keeps executing from the beginning as long as it is within the specified limit. If the limit for the inner loop is exceeded, the outer loop executes again, and so on.