## How Arrays and Slices work in Go: Why you need to create slice copies.

### Code extract:

```go
// Create copies of the slices to avoid modifying the same underlying array
		left := append([]int(nil), arr[:mid]...)
		right := append([]int(nil), arr[mid:]...)
```

In Go, slices are references to arrays. When you create a slice from an array, the slice still points to the same underlying array. Any modifications to the slice will affect the original array. This can cause problems in recursive algorithms like Merge Sort.

Let's take a case-study

#### Example Array
```go
arr := []int{38, 27, 43, 3, 9, 82, 10}
```

### Initial Split

1. **First Split:**
   ```go
   mid := len(arr) / 2
   left := arr[:mid] // left = [38, 27, 43]
   right := arr[mid:] // right = [3, 9, 82, 10]
   ```
   - `left` points to `arr[0:3]` and `right` points to `arr[3:7]`.

### Recursive Sorting and Modification

2. **Recursive Call for `left`:**
   ```go
   mergeSort(left) // Sorts [38, 27, 43]
   ```
   - Let's assume `left` is sorted to `[27, 38, 43]`.
   - `arr` now looks like `[27, 38, 43, 3, 9, 82, 10]` because `left` is a view into `arr`.

3. **Recursive Call for `right`:**
   ```go
   mergeSort(right) // Sorts [3, 9, 82, 10]
   ```
   - After sorting, `right` becomes `[3, 9, 10, 82]`.
   - `arr` now looks like `[27, 38, 43, 3, 9, 10, 82]` because `right` is a view into `arr`.

### Merging Process

4. **Merge Sorted Halves:**
   ```go
   i, j, k := 0, 0, 0
   for i < len(left) && j < len(right) {
       if left[i] < right[j] {
           arr[k] = left[i]
           i++
       } else {
           arr[k] = right[j]
           j++
       }
       k++
   }
   ```
   - `arr[k]` is being updated in place based on the comparison between `left[i]` and `right[j]`.

### Potential Problems Without Copying Slices

- **In-Place Modification:** The problem arises because `left` and `right` are slices of `arr`. When you modify `left` and `right` (by sorting and merging them), you are modifying the same array `arr` they are slices of, potentially leading to an unintended output.

- **Recursive Overwrites:** If the slices overlap or are manipulated incorrectly, the sorting and merging logic can overwrite parts of the array or produce incorrect results due to unintended side effects.

### Correct Approach: Copying Slices

To avoid these issues, make sure that each slice created during sorting is independent of the original array and any other slices. Here’s how you handle this:

1. **Copy the Slice Data:**
   ```go
   left := append([]int(nil), arr[:mid]...) // Creates a new slice with the data of arr[:mid]
   right := append([]int(nil), arr[mid:]...) // Creates a new slice with the data of arr[mid:]
   ```

2. **Recursive Calls:**
   ```go
   mergeSort(left)  // Operates on a new slice
   mergeSort(right) // Operates on a new slice
   ```

3. **Merging:**
   - After sorting, merge `left` and `right` back into `arr`:
   ```go
   i, j, k := 0, 0, 0
   for i < len(left) && j < len(right) {
       if left[i] < right[j] {
           arr[k] = left[i]
           i++
       } else {
           arr[k] = right[j]
           j++
       }
       k++
   }

   // Copy remaining elements from left
   for i < len(left) {
       arr[k] = left[i]
       i++
       k++
   }

   // Copy remaining elements from right
   for j < len(right) {
       arr[k] = right[j]
       j++
       k++
   }
   ```

   **NB:** In Go, the ellipses(`...`) are used in the context of slices to indicate that a slice should be expanded or unpacked into individual elements. When used with the `append` function, it means that all elements from the slice should be appended individually to the target slice (as opposed to appending them as a slice)

#### Visual example

Suppose result is [1, 2], and left[i:] is [3, 4, 5]:

- **Without Ellipses:** append(result, left[i:]) would give [1, 2, [3, 4, 5]]

- **With Ellipses:** append(result, left[i:]...) would give [1, 2, 3, 4, 5] (unpacking the slice and adding each element individually).

In conclusion, the ellipses are a way to "unpack" the slice into its individual elements when appending.