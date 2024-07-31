# Code Simulation

Here is a staged example of the `lengthOfLongestSubstring`. The function basically finds the length of the longest substring without repeating characters.

Consider the input string `s = "abcabcbb"`.

#### **Step-by-Step Analysis**

1. **Initialization**:
   - `charIndex` = `{}` (map to store the last index of each character)
   - `start` = 0 (beginning of the current window)
   - `maxLength` = 0 (length of the longest substring found)

2. **Iteration**:

   **Iteration 1** (`end` = 0):
   - `s[end]` = `'a'`
   - `charIndex` does not contain `'a'`, so no update to `start`.
   - Update `charIndex`: `{'a': 0}`
   - Calculate `maxLength`: `max(0, 0 - 0 + 1) = 1`
   - `maxLength` = 1

   **Iteration 2** (`end` = 1):
   - `s[end]` = `'b'`
   - `charIndex` does not contain `'b'`, so no update to `start`.
   - Update `charIndex`: `{'a': 0, 'b': 1}`
   - Calculate `maxLength`: `max(1, 1 - 0 + 1) = 2`
   - `maxLength` = 2

   **Iteration 3** (`end` = 2):
   - `s[end]` = `'c'`
   - `charIndex` does not contain `'c'`, so no update to `start`.
   - Update `charIndex`: `{'a': 0, 'b': 1, 'c': 2}`
   - Calculate `maxLength`: `max(2, 2 - 0 + 1) = 3`
   - `maxLength` = 3

   **Iteration 4** (`end` = 3):
   - `s[end]` = `'a'`
   - `charIndex` contains `'a'` at index 0 which is >= `start` (0), so update `start` to `1`.
   - Update `charIndex`: `{'a': 3, 'b': 1, 'c': 2}`
   - Calculate `maxLength`: `max(3, 3 - 1 + 1) = 3`
   - `maxLength` remains 3

   **Iteration 5** (`end` = 4):
   - `s[end]` = `'b'`
   - `charIndex` contains `'b'` at index 1 which is >= `start` (1), so update `start` to `2`.
   - Update `charIndex`: `{'a': 3, 'b': 4, 'c': 2}`
   - Calculate `maxLength`: `max(3, 4 - 2 + 1) = 3`
   - `maxLength` remains 3

   **Iteration 6** (`end` = 5):
   - `s[end]` = `'c'`
   - `charIndex` contains `'c'` at index 2 which is >= `start` (2), so update `start` to `3`.
   - Update `charIndex`: `{'a': 3, 'b': 4, 'c': 5}`
   - Calculate `maxLength`: `max(3, 5 - 3 + 1) = 3`
   - `maxLength` remains 3

   **Iteration 7** (`end` = 6):
   - `s[end]` = `'b'`
   - `charIndex` contains `'b'` at index 4 which is >= `start` (3), so update `start` to `5`.
   - Update `charIndex`: `{'a': 3, 'b': 6, 'c': 5}`
   - Calculate `maxLength`: `max(3, 6 - 5 + 1) = 3`
   - `maxLength` remains 3

   **Iteration 8** (`end` = 7):
   - `s[end]` = `'b'`
   - `charIndex` contains `'b'` at index 6 which is >= `start` (5), so update `start` to `7`.
   - Update `charIndex`: `{'a': 3, 'b': 7, 'c': 5}`
   - Calculate `maxLength`: `max(3, 7 - 7 + 1) = 3`
   - `maxLength` remains 3

3. **Result**:
   - After processing all characters, `maxLength` is 3.

### Summary

The longest substring without repeating characters in `"abcabcbb"` is `"abc"`, which has a length of 3. The `lengthOfLongestSubstring` function efficiently computes this by maintaining a sliding window and dynamically updating the start position and length based on character occurrences.

This approach ensures that each character is processed once, resulting in a time complexity of O(n), where n is the length of the string.