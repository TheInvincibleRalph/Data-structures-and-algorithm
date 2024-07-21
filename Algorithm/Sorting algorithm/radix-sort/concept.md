
Given the array:
```
[38, 27, 43, 3, 9, 82, 10]
```

## Step-by-Step Simulation on Ones Place Value

### Counting Occurrences by Ones Place
This process sorts the array by the least significant digit. For example, if 7 is in one-th place in a digit, it is placed at index 7 in a counting array.

1. **Count Array Initialization:**
   ```
   [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
   ```

2. **Number: 38, Ones Digit: 8 -> Count[8]++**
   ```
   [0, 0, 0, 0, 0, 0, 0, 0, 1, 0]
   ```

3. **Number: 27, Ones Digit: 7 -> Count[7]++**
   ```
   [0, 0, 0, 0, 0, 0, 0, 1, 1, 0]
   ```

4. **Number: 43, Ones Digit: 3 -> Count[3]++**
   ```
   [0, 0, 0, 1, 0, 0, 0, 1, 1, 0]
   ```

5. **Number: 3, Ones Digit: 3 -> Count[3]++**
   ```
   [0, 0, 0, 2, 0, 0, 0, 1, 1, 0]
   ```

6. **Number: 9, Ones Digit: 9 -> Count[9]++**
   ```
   [0, 0, 0, 2, 0, 0, 0, 1, 1, 1]
   ```

7. **Number: 82, Ones Digit: 2 -> Count[2]++**
   ```
   [0, 0, 1, 2, 0, 0, 0, 1, 1, 1]
   ```

8. **Number: 10, Ones Digit: 0 -> Count[0]++**
   ```
   [1, 0, 1, 2, 0, 0, 0, 1, 1, 1]
   ```

### Transforming to Cumulative Count Array

1. **Count[1] = Count[1] + Count[0] = 0 + 1 = 1**
   ```
   [1, 1, 1, 2, 0, 0, 0, 1, 1, 1]
   ```

2. **Count[2] = Count[2] + Count[1] = 1 + 1 = 2**
   ```
   [1, 1, 2, 2, 0, 0, 0, 1, 1, 1]
   ```

3. **Count[3] = Count[3] + Count[2] = 2 + 2 = 4**
   ```
   [1, 1, 2, 4, 0, 0, 0, 1, 1, 1]
   ```

4. **Count[4] = Count[4] + Count[3] = 0 + 4 = 4**
   ```
   [1, 1, 2, 4, 4, 0, 0, 1, 1, 1]
   ```

5. **Count[5] = Count[5] + Count[4] = 0 + 4 = 4**
   ```
   [1, 1, 2, 4, 4, 4, 0, 1, 1, 1]
   ```

6. **Count[6] = Count[6] + Count[5] = 0 + 4 = 4**
   ```
   [1, 1, 2, 4, 4, 4, 4, 1, 1, 1]
   ```

7. **Count[7] = Count[7] + Count[6] = 1 + 4 = 5**
   ```
   [1, 1, 2, 4, 4, 4, 4, 5, 1, 1]
   ```

8. **Count[8] = Count[8] + Count[7] = 1 + 5 = 6**
   ```
   [1, 1, 2, 4, 4, 4, 4, 5, 6, 1]
   ```

9. **Count[9] = Count[9] + Count[8] = 1 + 6 = 7**
   ```
   [1, 1, 2, 4, 4, 4, 4, 5, 6, 7]
   ```

### Building the Output Array

1. **Output Array:**
   ```
   [0, 0, 0, 0, 0, 0, 0]
   ```

2. **Place Number: 10, Ones Digit: 0**
   - Place 10 at index `count[0] - 1 = 0`
   - Decrement `count[0]`
   ```
   Output Array: [10, 0, 0, 0, 0, 0, 0]
   Count Array: [0, 1, 2, 4, 4, 4, 4, 5, 6, 7]
   ```

3. **Place Number: 82, Ones Digit: 2**
   - Place 82 at index `count[2] - 1 = 1`
   - Decrement `count[2]`
   ```
   Output Array: [10, 82, 0, 0, 0, 0, 0]
   Count Array: [0, 1, 1, 4, 4, 4, 4, 5, 6, 7]
   ```

4. **Place Number: 9, Ones Digit: 9**
   - Place 9 at index `count[9] - 1 = 6`
   - Decrement `count[9]`
   ```
   Output Array: [10, 82, 0, 0, 0, 0, 9]
   Count Array: [0, 1, 1, 4, 4, 4, 4, 5, 6, 6]
   ```

5. **Place Number: 3, Ones Digit: 3**
   - Place 3 at index `count[3] - 1 = 3`
   - Decrement `count[3]`
   ```
   Output Array: [10, 82, 0, 3, 0, 0, 9]
   Count Array: [0, 1, 1, 3, 4, 4, 4, 5, 6, 6]
   ```

6. **Place Number: 43, Ones Digit: 3**
   - Place 43 at index `count[3] - 1 = 2`
   - Decrement `count[3]`
   ```
   Output Array: [10, 82, 43, 3, 0, 0, 9]
   Count Array: [0, 1, 1, 2, 4, 4, 4, 5, 6, 6]
   ```

7. **Place Number: 27, Ones Digit: 7**
   - Place 27 at index `count[7] - 1 = 4`
   - Decrement `count[7]`
   ```
   Output Array: [10, 82, 43, 3, 27, 0, 9]
   Count Array: [0, 1, 1, 2, 4, 4, 4, 4, 6, 6]
   ```

8. **Place Number: 38, Ones Digit: 8**
   - Place 38 at index `count[8] - 1 = 5`
   - Decrement `count[8]`
   ```
   Output Array: [10, 82, 43, 3, 27, 38, 9]
   Count Array: [0, 1, 1, 2, 4, 4, 4, 4, 5, 6]
   ```

### Sorted by Ones Place

The array sorted by the ones place is:

```
[10, 82, 43, 3, 27, 38, 9]
```

## Step-by-Step Simulation on Tens Place Value

Given the array after sorting by the least significant digit:
```
[10, 82, 3, 43, 27, 38, 9]
```

### Counting Occurrences by Tens Place
Let's start by counting the occurrences of each tens digit.

1. **Count Array Initialization:**
   ```
   [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
   ```

2. **Number: 10, Tens Digit: 1 -> Count[1]++**
   ```
   [0, 1, 0, 0, 0, 0, 0, 0, 0, 0]
   ```

3. **Number: 82, Tens Digit: 8 -> Count[8]++**
   ```
   [0, 1, 0, 0, 0, 0, 0, 0, 1, 0]
   ```

4. **Number: 3, Tens Digit: 0 -> Count[0]++**
   ```
   [1, 1, 0, 0, 0, 0, 0, 0, 1, 0]
   ```

5. **Number: 43, Tens Digit: 4 -> Count[4]++**
   ```
   [1, 1, 0, 0, 1, 0, 0, 0, 1, 0]
   ```

6. **Number: 27, Tens Digit: 2 -> Count[2]++**
   ```
   [1, 1, 1, 0, 1, 0, 0, 0, 1, 0]
   ```

7. **Number: 38, Tens Digit: 3 -> Count[3]++**
   ```
   [1, 1, 1, 1, 1, 0, 0, 0, 1, 0]
   ```

8. **Number: 9, Tens Digit: 0 -> Count[0]++**
   ```
   [2, 1, 1, 1, 1, 0, 0, 0, 1, 0]
   ```

### Transforming to Cumulative Count Array
1. **Count[1] = Count[1] + Count[0] = 1 + 2 = 3**
   ```
   [2, 3, 1, 1, 1, 0, 0, 0, 1, 0]
   ```

2. **Count[2] = Count[2] + Count[1] = 1 + 3 = 4**
   ```
   [2, 3, 4, 1, 1, 0, 0, 0, 1, 0]
   ```

3. **Count[3] = Count[3] + Count[2] = 1 + 4 = 5**
   ```
   [2, 3, 4, 5, 1, 0, 0, 0, 1, 0]
   ```

4. **Count[4] = Count[4] + Count[3] = 1 + 5 = 6**
   ```
   [2, 3, 4, 5, 6, 0, 0, 0, 1, 0]
   ```

5. **Count[5] = Count[5] + Count[4] = 0 + 6 = 6**
   ```
   [2, 3, 4, 5, 6, 6, 0, 0, 1, 0]
   ```

6. **Count[6] = Count[6] + Count[5] = 0 + 6 = 6**
   ```
   [2, 3, 4, 5, 6, 6, 6, 0, 1, 0]
   ```

7. **Count[7] = Count[7] + Count[6] = 0 + 6 = 6**
   ```
   [2, 3, 4, 5, 6, 6, 6, 6, 1, 0]
   ```

8. **Count[8] = Count[8] + Count[7] = 1 + 6 = 7**
   ```
   [2, 3, 4, 5, 6, 6, 6, 6, 7, 0]
   ```

9. **Count[9] = Count[9] + Count[8] = 0 + 7 = 7**
   ```
   [2, 3, 4, 5, 6, 6, 6, 6, 7, 7]
   ```

### Building the Output Array

1. **Output Array:**
   ```
   [0, 0, 0, 0, 0, 0, 0]
   ```

2. **Place Number: 9, Tens Digit: 0**
   - Place 9 at index `count[0] - 1 = 1`
   - Decrement `count[0]`
   ```
   Output Array: [0, 9, 0, 0, 0, 0, 0]
   Count Array: [1, 3, 4, 5, 6, 6, 6, 6, 7, 7]
   ```

3. **Place Number: 38, Tens Digit: 3**
   - Place 38 at index `count[3] - 1 = 4`
   - Decrement `count[3]`
   ```
   Output Array: [0, 9, 0, 0, 38, 0, 0]
   Count Array: [1, 3, 4, 4, 6, 6, 6, 6, 7, 7]
   ```

4. **Place Number: 27, Tens Digit: 2**
   - Place 27 at index `count[2] - 1 = 3`
   - Decrement `count[2]`
   ```
   Output Array: [0, 9, 0, 27, 38, 0, 0]
   Count Array: [1, 3, 3, 4, 6, 6, 6, 6, 7, 7]
   ```

5. **Place Number: 43, Tens Digit: 4**
   - Place 43 at index `count[4] - 1 = 5`
   - Decrement `count[4]`
   ```
   Output Array: [0, 9, 0, 27, 38, 43, 0]
   Count Array: [1, 3, 3, 4, 5, 6, 6, 6, 7, 7]
   ```

6. **Place Number: 3, Tens Digit: 0**
   - Place 3 at index `count[0] - 1 = 0`
   - Decrement `count[0]`
   ```
   Output Array: [3, 9, 0, 27, 38, 43, 0]
   Count Array: [0, 3, 3, 4, 5, 6, 6, 6, 7, 7]
   ```

7. **Place Number: 82, Tens Digit: 8**
   - Place 82 at index `count[8] - 1 = 6`
   - Decrement `count[8]`
   ```
   Output Array: [3, 9, 0, 27, 38, 43, 82]
   Count Array: [0, 3, 3, 4, 5, 6, 6, 6, 6, 7]
   ```

8. **Place Number: 10, Tens Digit: 1**
   - Place 10 at index `count[1] - 1 = 2`
   - Decrement `count[1]`
   ```
   Output Array: [3, 9, 10, 27, 38, 43, 82]
   Count Array: [0, 2, 3, 4, 5, 6, 6, 6, 6, 7]
   ```

### Correctly Sorted Array by Tens Place

The final output, which is the array sorted by the tens place, should now be:
```
[3, 9, 10, 27, 38, 43, 82]
```

## Logic Behind the Code

1. **The Count Array:**
   - This takes record of the number of times each digit appears at the current place value.

2. **Cumulative Count:**
   - This process updates the `count` array to reflect cumulative count of the occurences. In this step, each position in the `count` array will store the sum of counts up to that position.
   - The cumulative count helps determine the exact index at which each element should be placed in the output array.

3. **The Output Array:**
   - The cumulative count array helps to place each element in the correct position in the output array.

   - The output array at the end of each place value operation is first traversed from right to left (to maintain stability-the relative order of elements with equal values).

   - The cumulative count array is then continuously adjusted as elements are placed in the output array.

### Detailed Example

Consider sorting by the ones place with the array:
```
[38, 27, 43, 3, 9, 82, 10]
```

### Step 1: Counting Occurrences

1. Initialize a count array for digit values (0-9):
   ```
   count = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
   ```

2. Count the occurrences of each ones digit:
   ```
   Number -> one-th value => count increment

         38 -> 8  => count[8]++
         27 -> 7  => count[7]++
         43 -> 3  => count[3]++
         3  -> 3  => count[3]++
         9  -> 9  => count[9]++
         82 -> 2  => count[2]++
         10 -> 0  => count[0]++
   ```

   Resulting count array:
   ```
   count = [1, 0, 1, 2, 0, 0, 0, 1, 1, 1]
   ```

### Step 2: Transform to Cumulative Count

Update the count array to reflect cumulative counts:
1. Add each count to the count of the previous index:
   ```
   Initial count = initial count + preceeding digit 

      count[1] = count[1] + count[0] => 0 + 1 = 1
      count[2] = count[2] + count[1] => 1 + 1 = 2
      count[3] = count[3] + count[2] => 2 + 2 = 4
      count[4] = count[4] + count[3] => 0 + 4 = 4
      count[5] = count[5] + count[4] => 0 + 4 = 4
      count[6] = count[6] + count[5] => 0 + 4 = 4
      count[7] = count[7] + count[6] => 1 + 4 = 5
      count[8] = count[8] + count[7] => 1 + 5 = 6
      count[9] = count[9] + count[8] => 1 + 6 = 7
   ```

   Resulting cumulative count array:
   ```
   count = [1, 1, 2, 4, 4, 4, 4, 5, 6, 7]
   ```

### Step 3: Building the Output Array

1. Initialize an output array of the same length:
   ```
   output = [0, 0, 0, 0, 0, 0, 0]
   ```

2. Traverse the original array from right to left, placing each element in its correct position using the cumulative count array, and then decrement the cumulative count for that digit:

   - **Number: 10, Ones Digit: 0**
     - Place 10 at `count[0] - 1 = 0`
     - Decrement `count[0]`
     ```
     output = [10, 0, 0, 0, 0, 0, 0]
     count = [0, 1, 2, 4, 4, 4, 4, 5, 6, 7]
     ```

   - **Number: 82, Ones Digit: 2**
     - Place 82 at `count[2] - 1 = 1`
     - Decrement `count[2]`
     ```
     output = [10, 82, 0, 0, 0, 0, 0]
     count = [0, 1, 1, 4, 4, 4, 4, 5, 6, 7]
     ```

   - **Number: 9, Ones Digit: 9**
     - Place 9 at `count[9] - 1 = 6`
     - Decrement `count[9]`
     ```
     output = [10, 82, 0, 0, 0, 0, 9]
     count = [0, 1, 1, 4, 4, 4, 4, 5, 6, 6]
     ```

   - **Number: 3, Ones Digit: 3**
     - Place 3 at `count[3] - 1 = 3`
     - Decrement `count[3]`
     ```
     output = [10, 82, 0, 3, 0, 0, 9]
     count = [0, 1, 1, 3, 4, 4, 4, 5, 6, 6]
     ```

   - **Number: 43, Ones Digit: 3**
     - Place 43 at `count[3] - 1 = 2`
     - Decrement `count[3]`
     ```
     output = [10, 82, 43, 3, 0, 0, 9]
     count = [0, 1, 1, 2, 4, 4, 4, 5, 6, 6]
     ```

   - **Number: 27, Ones Digit: 7**
     - Place 27 at `count[7] - 1 = 4`
     - Decrement `count[7]`
     ```
     output = [10, 82, 43, 3, 27, 0, 9]
     count = [0, 1, 1, 2, 4, 4, 4, 4, 6, 6]
     ```

   - **Number: 38, Ones Digit: 8**
     - Place 38 at `count[8] - 1 = 5`
     - Decrement `count[8]`
     ```
     output = [10, 82, 43, 3, 27, 38, 9]
     count = [0, 1, 1, 2, 4, 4, 4, 4, 5, 6]
     ```

### Result

After sorting by the ones place, the array is:
```
[10, 82, 43, 3, 27, 38, 9]
```


















