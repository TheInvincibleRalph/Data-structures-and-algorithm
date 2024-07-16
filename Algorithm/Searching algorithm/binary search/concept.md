## What is Binary Search?

In computer science, binary search, also known as half-interval search, logarithmic search, or binary chop, is a search algorithm that finds the position of a target value within a sorted array. Binary search compares the target value to the middle element of the array. If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half, again taking the middle element to compare to the target value, and repeating this until the target value is found. If the search ends with the remaining half being empty, the target is not in the array. [Wikipedia](https://en.wikipedia.org/wiki/Binary_search)



## How does it work?

In its simplest form, Binary Search operates on a contiguous sequence with a specified left and right index. This is called the Search Space. Binary Search maintains the left, right, and middle indicies of the search space and compares the search target or applies the search condition to the middle value of the collection; if the condition is unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful. If the search ends with an empty half, the condition cannot be fulfilled and target is not found. [Leetcode](https://leetcode.com/explore/learn/card/binary-search/138/background/971/)

#### References

In the sqRoot function: That particular formula is used because when using (left + right) / 2, if left and right are large, their sum left + right could exceed the maximum value that can be stored in an integer variable, leading to an overflow. On the other hand, left + (right - left) / 2 ensures that no overflow occurs because right - left is calculated first, which is always a non-negative value and much smaller than left + right