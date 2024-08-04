### Circular Queue

A circular queue is a linear data structure that follows the FIFO (First In, First Out) principle but connects the end of the queue back to the beginning, making it circular. 

### Key Properties

- **Front**: Points to the first element in the queue.
- **Rear**: Points to the last element in the queue.
- **Capacity**: The maximum number of elements the queue can hold.

### Full and Empty Conditions

- **Empty Condition**: The queue is empty when `front == rear`.
- **Full Condition**: The queue is full when `(rear + 1) % capacity == front`.

### Adding and Removing Elements

When an element is added (enqueue operation) to a circular queue:
- The `rear` index is incremented circularly.
- If the queue was `[1, 2, 3, _, _]` and you add `4`, the queue becomes `[1, 2, 3, 4, _]` and `rear` points to the next position.

When an element is removed (dequeue operation) from a circular queue:
- The `front` index is incremented circularly.
- If the queue was `[_, _, 3, 4, _]` and you remove `3`, the queue becomes `[_, _, _, 4, _]` and `front` points to the next position.

### Example

Here's a step-by-step simulation with a circular queue of capacity 5.

1. **Initial State**:
   - Queue: `[_ , _, _, _, _]`
   - `front = 0`
   - `rear = 0`

2. **Enqueue 1, 2, 3**:
   - Queue: `[1, 2, 3, _, _]`
   - `front = 0`
   - `rear = 3`

3. **Dequeue 1, 2**:
   - Queue: `[_, _, 3, _, _]`
   - `front = 2`
   - `rear = 3`

4. **Enqueue 4**:
   - Queue: `[_, _, 3, 4, _]`
   - `front = 2`
   - `rear = 4`

5. **Enqueue 5**:
   - Queue: `[_, _, 3, 4, 5]`
   - `front = 2`
   - `rear = 0` (wraps around)