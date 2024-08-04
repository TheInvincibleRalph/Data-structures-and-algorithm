## Circular Queue

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

 ## Simulation of the Full Condition

1. **Initialization**:
   - Create a circular queue with a capacity of 5. The underlying array, however, should have 6 slots. 
   - `front` and `rear` both start at 0.

2. **Enqueue Operations**:
   - Enqueue 1: Items = `[1, _, _, _, _, _]`, `front = 0`, `rear = 1`
   - Enqueue 2: Items = `[1, 2, _, _, _, _]`, `front = 0`, `rear = 2`
   - Enqueue 3: Items = `[1, 2, 3, _, _, _]`, `front = 0`, `rear = 3`
   - Enqueue 4: Items = `[1, 2, 3, 4, _, _]`, `front = 0`, `rear = 4`
   - Enqueue 5: Items = `[1, 2, 3, 4, 5, _]`, `front = 0`, `rear = 5`

3. **Full Queue Condition**:
   - When `rear` is at index 5, adding 1 (to check full condition) gives 6, and `(6) % 6 == 0` which equals `front`. Therefore, the queue is full.
   - The output of `Is Full: true` confirms that the queue is full.

4. **Further Enqueue**:
   - Attempting to enqueue 6 fails with the output `Queue is full` because the queue is indeed full.

5. **Dequeue Operations**:
   - Dequeueing elements one by one will show `[1, 2, 3, 4, 5]`.

6. **Empty Queue Condition**:
   - After dequeueing all elements, both `front` and `rear` will be 0, confirming the queue is empty with `IsEmpty: true`.

### Extra Points

- The extra slot in the array helps in distinguishing between full and empty states.
- The full condition `(rear + 1) % capacity == front` is used to confirm the queue is full.

## The Empty Condition `front == rear`

- The `front` points to the position of the first element in the queue (the next one to be dequeued).
- The `rear` points to the position just after the last element in the queue (where the next element will be enqueued)

### Why the condition?

The front pointer shows the position of the next element to be dequeued, and if it is the same as the rear pointer (which points to the space where the next element will be enqueued), then there are no elements to dequeue and by extension, the queue is empty.

### Extra points:
- The front position is always one position ahead of the rear position in a non-empty circular queue.
- Since the queue is circular, the rear position can be at the beginning of the queue when the end of the queue is reached (this is called wrapping to zero).
- The front and rear pointers cannot be pointing to the same position, if that happens, the queue is said to be empty.