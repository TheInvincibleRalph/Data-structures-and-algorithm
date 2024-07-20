### Data Structure: Binary Tree

A binary tree is a hierarchical data structure in which each node has at most two children, referred to as the left child and the right child.

### Key Characteristics of a Binary Tree

1. **Node:**
   - Each element in the tree is called a node.
   - A node contains a value and pointers (references) to its children.

2. **Root:**
   - The topmost node in a binary tree.
   - There is only one root node in a binary tree.

3. **Children:**
   - The nodes directly connected to a node are called its children.
   - A node can have zero, one, or two children in a binary tree.

4. **Parent:**
   - The node directly above a node is called its parent.
   - Each node (except the root) has one parent.

5. **Leaf:**
   - A node with no children is called a leaf node or terminal node.

6. **Subtree:**
   - A subtree is a tree consisting of a node and its descendants.

### Binary Heap

Binary heap is a *special case* of binary tree

A binary heap is a complete binary tree which satisfies the heap property.

There are two types of binary heaps:
- **Max Heap:** The value of each node is greater than or equal to the values of its children.
- **Min Heap:** The value of each node is less than or equal to the values of its children.

### Binary Heap in Heap Sort Algorithm

Heap sort utilizes a binary heap to sort an array. 

Here’s how the concepts of a binary tree relate to the heap sort algorithm:

#### Heap Sort Code Revisited

```go
func heapSort(arr []int) {
    n := len(arr)
    
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    
    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        
        heapify(arr, i, 0)
    }
}

func heapify(arr []int, n, i int) {
    largest := i       // Initialize largest as root
    left := 2*i + 1    // left = 2*i + 1
    right := 2*i + 2   // right = 2*i + 2

    
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i] 

        heapify(arr, n, largest)
    }
}
```

#### Relation to Binary Tree

1. **Array Representation:**
   - A binary heap is often represented as an array. For any node at index `i`:
     - Its left child is at index `2*i + 1`.
     - Its right child is at index `2*i + 2`.
   - This property allows easy navigation between parent and children nodes.

2. **Building the Heap:**
   - The `for` loop in `heapSort` builds the heap by calling `heapify` on all non-leaf nodes starting from the last non-leaf node (`n/2 - 1`) up to the root.
   - This ensures that each subtree satisfies the heap property.

3. **Heapify Process:**
   - `heapify` ensures that the subtree rooted at index `i` satisfies the max heap property.
   - It compares the node at `i` with its left and right children.
   - If necessary, it swaps the node with the larger of its children and recursively heapifies the affected subtree.

4. **Sorting:**
   - After building the max heap, the largest element (root) is at `arr[0]`.
   - The `for` loop then repeatedly extracts the largest element by swapping the root with the last element of the heap and reducing the heap size by one.
   - `heapify` is called on the reduced heap to restore the max heap property.

#### Example

Consider the array `[38, 27, 43, 3, 9, 82, 10]`:

1. **Building the Max Heap:**
   - Transform the array into a max heap.
   - Array after building max heap: `[82, 27, 43, 3, 9, 38, 10]`.

2. **Sorting:**
   - Extract the largest element (82) and move it to the end.
   - Heapify the reduced array `[10, 27, 43, 3, 9, 38]`.
   - Repeat the process until the array is sorted.

### Concept Simulation


Let's perform a heap sort simulation on the array: `[38, 27, 43, 3, 9, 82, 10]`.

### Step 1: Build an initial Max Heap

First, we need to build a max heap from the array. A max heap is a binary tree where each parent node is greater than or equal to its child nodes.

1. **Heapify the sub-tree rooted at index 2:**

   - Array: `[38, 27, 43, 3, 9, 82, 10]`
   - Children of 43 (index 2): 82 (index 5) and 10 (index 6)
   - 82 is larger than 43, so we swap them:
   - Result: `[38, 27, 82, 3, 9, 43, 10]`

2. **Heapify the sub-tree rooted at index 1:**

   - Array: `[38, 27, 82, 3, 9, 43, 10]`
   - Children of 27 (index 1): 3 (index 3) and 9 (index 4)
   - 27 is larger than both children, so no swap is needed:
   - Result: `[38, 27, 82, 3, 9, 43, 10]`

3. **Heapify the sub-tree rooted at index 0:**

   - Array: `[38, 27, 82, 3, 9, 43, 10]`
   - Children of 38 (index 0): 27 (index 1) and 82 (index 2)
   - 82 is larger than 38, so we swap them:
   - Result: `[82, 27, 38, 3, 9, 43, 10]`
   - Now, heapify the sub-tree rooted at index 2:
   - Children of 38 (index 2): 43 (index 5) and 10 (index 6)
   - 43 is larger than 38, so we swap them:
   - Result: `[82, 27, 43, 3, 9, 38, 10]`

At this point, we have built a max heap: `[82, 27, 43, 3, 9, 38, 10]`.

          82
        /    \
       43     38
      / \     / \
    27   3   9   10
*Max Heap Representation*

### Step 2: Extract Elements from the Heap

Now we will extract the largest element (the root of the heap) and move it to the end of the array. We then rebuild the heap and repeat the process.

1. **Move the root (82) to the end:**

   - Swap 82 with 10:
   - Array: `[10, 27, 43, 3, 9, 38, 82]`
   - Heapify the reduced heap of size 6:
   - Root (10), Children: 27 (index 1) and 43 (index 2)
   - 43 is the largest, so swap 10 with 43:
   - Result: `[43, 27, 10, 3, 9, 38, 82]`
   - Heapify sub-tree rooted at index 2:
   - Children: 38 (index 5)
   - 38 is larger than 10, so swap them:
   - Result: `[43, 27, 38, 3, 9, 10, 82]`

2. **Move the root (43) to the end:**

   - Swap 43 with 10:
   - Array: `[10, 27, 38, 3, 9, 43, 82]`
   - Heapify the reduced heap of size 5:
   - Root (10), Children: 27 (index 1) and 38 (index 2)
   - 38 is the largest, so swap 10 with 38:
   - Result: `[38, 27, 10, 3, 9, 43, 82]`
   - No need to heapify sub-tree at index 2 as it has no children

3. **Move the root (38) to the end:**

   - Swap 38 with 9:
   - Array: `[9, 27, 10, 3, 38, 43, 82]`
   - Heapify the reduced heap of size 4:
   - Root (9), Children: 27 (index 1) and 10 (index 2)
   - 27 is the largest, so swap 9 with 27:
   - Result: `[27, 9, 10, 3, 38, 43, 82]`
   - No need to heapify sub-tree at index 1 as it has no children

4. **Move the root (27) to the end:**

   - Swap 27 with 3:
   - Array: `[3, 9, 10, 27, 38, 43, 82]`
   - Heapify the reduced heap of size 3:
   - Root (3), Children: 9 (index 1) and 10 (index 2)
   - 10 is the largest, so swap 3 with 10:
   - Result: `[10, 9, 3, 27, 38, 43, 82]`

5. **Move the root (10) to the end:**

   - Swap 10 with 3:
   - Array: `[3, 9, 10, 27, 38, 43, 82]`
   - Heapify the reduced heap of size 2:
   - Root (3), Child: 9 (index 1)
   - 9 is larger, so swap 3 with 9:
   - Result: `[9, 3, 10, 27, 38, 43, 82]`

6. **Move the root (9) to the end:**

   - Swap 9 with 3:
   - Array: `[3, 9, 10, 27, 38, 43, 82]`

Now the array is fully sorted: `[3, 9, 10, 27, 38, 43, 82]`.


