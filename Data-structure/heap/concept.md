# Heap Operations

### `heapifyUp`

**Purpose**: `heapifyUp` (also known as `sift up` or `bubble up`) is used to restore the heap property after an element is added to the heap. 

**Context**: When you insert a new element into a heap, it is initially placed at the end of the heap. This can violate the heap property, so you use `heapifyUp` to move the newly added element up the heap until the heap property is restored.

**Operation**:
1. Compare the newly inserted element with its parent.
2. If the new element violates the heap property (e.g., in a min-heap, if it is smaller than its parent), swap it with the parent.
3. Repeat this process until the heap property is satisfied or you reach the root of the heap.

### `heapifyDown`

**Purpose**: `heapifyDown` (also known as `sift down` or `bubble down`) is used to restore the heap property after the root element is removed or after a replacement operation, such as when the last element is moved to the root position.

**Context**: After removing the root element from the heap (or after replacing it), the heap property might be violated, so `heapifyDown` is used to move the element at the root down the heap until the heap property is restored.

**Operation**:
1. Compare the element at the root with its children.
2. If the element violates the heap property (e.g., in a min-heap, if it is larger than any of its children), swap it with the smaller child (in a min-heap).
3. Repeat this process until the heap property is satisfied or you reach a leaf node.

### Example

1. **Insertion (heapifyUp)**:
   - Insert element `10` into a min-heap: `[5, 7, 15, 20, 10]`
   - After insertion, `10` is at the end: `[5, 7, 15, 20, 10]`
   - `heapifyUp` would compare `10` with its parent (`15`), and since `10` is smaller, swap them.
   - Resulting heap: `[5, 7, 10, 20, 15]`

2. **Removal (heapifyDown)**:
   - Remove the root element (`5`): `[7, 10, 15, 20]`
   - Move the last element (`20`) to the root: `[20, 7, 15, 10]`
   - `heapifyDown` would compare `20` with its children (`7` and `15`) and swap it with the smaller child (`7`).
   - Resulting heap: `[7, 10, 15, 20]`