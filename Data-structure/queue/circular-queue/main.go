package main

//since it is a circular queue, the front position cannot be the same as the rear position, otherwise the queue is empty
//the reason is because the front position is always one position ahead of the rear position in a non-empty circular queue
//The front points to the position of the first element in the queue (the next one to be dequeued).
//The rear points to the position just after the last element in the queue (where the next element will be enqueued)
//since the queue is circular, the rear position can be at the beginning of the queue when the end of the queue is reached
//however, the front and rear pointers cannot be pointing to the same position, if that happens, the queue is said to be empty.
// The front pointer shows the position of the next element to be dequeued,
// and if it is the same as the rear pointer (which points to the space where the next element will be enqueued),
// hen there are no elements to dequeue and by extension, the queue is empty.
