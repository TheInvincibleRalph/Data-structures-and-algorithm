
### Insertion Sort Analysis

```go
package main

import "fmt"

func insertionSort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1

        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1 //decrements the value of j by 1; the same as j--
        }
        arr[j+1] = key
    }
}

func main() {
    arr := []int{12, 11, 13, 5, 6}
    insertionSort(arr)
    fmt.Println("Sorted array is:", arr)
}
```

### Case Study

We'll be using the array `[12, 11, 13, 5, 6]` as a case-study for our simulation.

### Concept Simulation

#### Initial State

`arr = [12, 11, 13, 5, 6]`

#### Step-by-Step Simulation

1. **First Iteration (`i = 1`)**
    - `key = arr[1] = 11`
    - `j = i - 1 = 0`

    **Inner Loop:**
    - `j >= 0 && arr[j] > key` translates to `0 >= 0 && 12 > 11` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[1] = arr[0]`
      ```
      arr = [12, 12, 13, 5, 6]
      ```
    - Decrement `j`: `j = 0 - 1 = -1`
    - Exit inner loop since `j < 0`

    **Insert Key:**
    - `arr[j+1] = key` translates to `arr[0] = 11`
      ```
      arr = [11, 12, 13, 5, 6]
      ```

2. **Second Iteration (`i = 2`)**
    - `key = arr[2] = 13`
    - `j = i - 1 = 1`

    **Inner Loop:**
    - `j >= 0 && arr[j] > key` translates to `1 >= 0 && 12 > 13` (false)
    - Exit inner loop immediately

    **Insert Key:**
    - `arr[j+1] = key` translates to `arr[2] = 13` (no change needed)
      ```
      arr = [11, 12, 13, 5, 6]
      ```

3. **Third Iteration (`i = 3`)**
    - `key = arr[3] = 5`
    - `j = i - 1 = 2`

    **Inner Loop:**
    - `j >= 0 && arr[j] > key` translates to `2 >= 0 && 13 > 5` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[3] = arr[2]`
      ```
      arr = [11, 12, 13, 13, 6]
      ```
    - Decrement `j`: `j = 2 - 1 = 1`
    - `j >= 0 && arr[j] > key` translates to `1 >= 0 && 12 > 5` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[2] = arr[1]`
      ```
      arr = [11, 12, 12, 13, 6]
      ```
    - Decrement `j`: `j = 1 - 1 = 0`
    - `j >= 0 && arr[j] > key` translates to `0 >= 0 && 11 > 5` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[1] = arr[0]`
      ```
      arr = [11, 11, 12, 13, 6]
      ```
    - Decrement `j`: `j = 0 - 1 = -1`
    - Exit inner loop since `j < 0`

    **Insert Key:**
    - `arr[j+1] = key` translates to `arr[0] = 5`
      ```
      arr = [5, 11, 12, 13, 6]
      ```

4. **Fourth Iteration (`i = 4`)**
    - `key = arr[4] = 6`
    - `j = i - 1 = 3`

    **Inner Loop:**
    - `j >= 0 && arr[j] > key` translates to `3 >= 0 && 13 > 6` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[4] = arr[3]`
      ```
      arr = [5, 11, 12, 13, 13]
      ```
    - Decrement `j`: `j = 3 - 1 = 2`
    - `j >= 0 && arr[j] > key` translates to `2 >= 0 && 12 > 6` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[3] = arr[2]`
      ```
      arr = [5, 11, 12, 12, 13]
      ```
    - Decrement `j`: `j = 2 - 1 = 1`
    - `j >= 0 && arr[j] > key` translates to `1 >= 0 && 11 > 6` (true)
    - Shift `arr[j]` to `arr[j+1]`: `arr[2] = arr[1]`
      ```
      arr = [5, 11, 11, 12, 13]
      ```
    - Decrement `j`: `j = 1 - 1 = 0`
    - Exit inner loop since `j < 0`

    **Insert Key:**
    - `arr[j+1] = key` translates to `arr[1] = 6`
      ```
      arr = [5, 6, 11, 12, 13]
      ```

### Final Sorted Array

The array is now sorted:
```
[5, 6, 11, 12, 13]
```

### Summary

1. The outer loop iterates through each element starting from the second element.
2. The `key` variable holds the value of the element to be inserted into the sorted portion of the array.
3. The inner loop shifts elements to the

right if they are greater than the `key`.
4. After finding the correct position, the `key` is inserted into the array.

The use of `j--` allows the inner loop to move backward through the sorted portion of the array, shifting elements as needed to make room for the `key`. This process ensures that the array remains sorted after each iteration of the outer loop.