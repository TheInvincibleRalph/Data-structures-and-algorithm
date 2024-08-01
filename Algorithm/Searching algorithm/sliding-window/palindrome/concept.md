# Longest Palindrome

The technique used in the `longestPalindrome` function is called **expanding around the center** because it identifies potential palindromic substrings by choosing a central point (or points as in the case of even length palindromes) and expanding outwards to the left and right to see how far the palindrome extends.

### Why It's Called "Expanding Around the Center":

1. **Central Starting Point**: Each potential palindrome is assumed to have a center. This center can be a single character (for odd-length palindromes) or a pair of characters (for even-length palindromes).

2. **Outward Expansion**: From the chosen center, the function checks whether the characters on both sides are equal, which is a necessary condition for a string to be a palindrome. If they are equal, the boundaries are expanded further outward to continue checking.

3. **Palindrome Identification**: If at any point the characters are not equal, the expansion stops. The segment identified from this process is a palindrome because it reads the same forward and backward.

### Key Characteristics of Expanding Around the Center:

- **Efficiency**: The method is efficient, with a time complexity of O(n^2), where n is the length of the string. This is because for each character or pair of characters, the expansion can potentially go up to the length of the string.
- **Simplicity**: The technique is straightforward to implement and understand. It simply involves checking character equality and adjusting pointers, with no additional data structures or preprocessing required.



