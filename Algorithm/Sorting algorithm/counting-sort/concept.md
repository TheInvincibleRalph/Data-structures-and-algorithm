

### Step-by-Step Simulation

#### Initial State:
```
arr = [4, 2, 2, 8, 3, 3, 1]
```

#### Step 1: Find the Maximum Value
```go
max := arr[0]
for _, val := range arr {
    if val > max {
        max = val
    }
}
// max = 8
```
- Iterating through the array to find the maximum value, which is 8.

#### Step 2: Create and Initialize the Count Array
```go
count := make([]int, max+1)
// count = [0, 0, 0, 0, 0, 0, 0, 0, 0]
```
- A count array of size `max+1` (9) is created and initialized to all zeros.

#### Step 3: Populate the Count Array
```go
for _, val := range arr {
    count[val]++
}
// count = [0, 1, 2, 2, 1, 0, 0, 0, 1]
```
- Iterating through the array, update the count array:
  - `4` appears once: `count[4]++`
  - `2` appears twice: `count[2]++`
  - `8` appears once: `count[8]++`
  - `3` appears twice: `count[3]++`
  - `1` appears once: `count[1]++`

#### Step 4: Transform Count Array to Cumulative Count Array
```go
for i := 1; i <= max; i++ {
    count[i] += count[i-1]
}
// count = [0, 1, 3, 5, 6, 6, 6, 6, 7]
```
- Transform the count array to hold the cumulative count:
  - `count[1] = count[1] + count[0] = 1 + 0 = 1`
  - `count[2] = count[2] + count[1] = 2 + 1 = 3`
  - `count[3] = count[3] + count[2] = 2 + 3 = 5`
  - `count[4] = count[4] + count[3] = 1 + 5 = 6`
  - `count[5] = count[5] + count[4] = 0 + 6 = 6`
  - `count[6] = count[6] + count[5] = 0 + 6 = 6`
  - `count[7] = count[7] + count[6] = 0 + 6 = 6`
  - `count[8] = count[8] + count[7] = 1 + 6 = 7`

#### Step 5: Create Output Array
```go
output := make([]int, len(arr))
// output = [0, 0, 0, 0, 0, 0, 0]
```
- Create an output array of the same length as the input array, initialized to zeros.

#### Step 6: Build the Output Array
```go
for i := len(arr) - 1; i >= 0; i-- {
    output[count[arr[i]]-1] = arr[i]
    count[arr[i]]--
}
```
- Traverse the input array from right to left, placing elements in the correct position in the output array and decrementing the count:
  - `arr[6] = 1`: `output[count[1]-1] = output[0] = 1`; decrement `count[1]` to 0
  - `arr[5] = 3`: `output[count[3]-1] = output[4] = 3`; decrement `count[3]` to 4
  - `arr[4] = 3`: `output[count[3]-1] = output[3] = 3`; decrement `count[3]` to 3
  - `arr[3] = 8`: `output[count[8]-1] = output[6] = 8`; decrement `count[8]` to 6
  - `arr[2] = 2`: `output[count[2]-1] = output[2] = 2`; decrement `count[2]` to 2
  - `arr[1] = 2`: `output[count[2]-1] = output[1] = 2`; decrement `count[2]` to 1
  - `arr[0] = 4`: `output[count[4]-1] = output[5] = 4`; decrement `count[4]` to 5

```go
  // output = [1, 2, 2, 3, 3, 4, 8]
  ```
  
#### Step 7: Copy the Output Array to the Original Array
```go
copy(arr, output)
// arr = [1, 2, 2, 3, 3, 4, 8]
```
- Copy the sorted output array back to the original array.

### Final Output
```
Sorted array is: [1, 2, 2, 3, 3, 4, 8]
```