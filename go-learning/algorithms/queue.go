package algorithms

// Queue implements a FIFO (First-In-First-Out) data structure.
// Uses a slice for dynamic sizing.
type Queue struct {
	items []int
}

// NewQueue creates and returns a new empty queue.
func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

// Enqueue adds an element to the back of the queue.
// Time Complexity: O(1) amortized
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Dequeue removes and returns the front element from the queue.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(n) for slice-based implementation
// Note: For better performance, use CircularQueue or implement with linked list
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

// Front returns the front element without removing it.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(1)
func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

// Back returns the back element without removing it.
// Time Complexity: O(1)
func (q *Queue) Back() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[len(q.items)-1], true
}

// IsEmpty returns true if the queue has no elements.
// Time Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
// Time Complexity: O(1)
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue.
// Time Complexity: O(1)
func (q *Queue) Clear() {
	q.items = []int{}
}

// ----------------------------------------------------------------------------
// CircularQueue - Fixed-size circular queue with O(1) operations
// ----------------------------------------------------------------------------

// CircularQueue implements a fixed-size circular queue.
// All operations are O(1).
type CircularQueue struct {
	items    []int
	front    int
	rear     int
	size     int
	capacity int
}

// NewCircularQueue creates a circular queue with the given capacity.
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]int, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

// Enqueue adds an element to the back of the circular queue.
// Returns false if the queue is full.
// Time Complexity: O(1)
func (cq *CircularQueue) Enqueue(val int) bool {
	if cq.IsFull() {
		return false
	}
	cq.rear = (cq.rear + 1) % cq.capacity
	cq.items[cq.rear] = val
	cq.size++
	return true
}

// Dequeue removes and returns the front element.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(1)
func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	val := cq.items[cq.front]
	cq.front = (cq.front + 1) % cq.capacity
	cq.size--
	return val, true
}

// Front returns the front element without removing it.
// Time Complexity: O(1)
func (cq *CircularQueue) Front() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	return cq.items[cq.front], true
}

// Rear returns the rear element without removing it.
// Time Complexity: O(1)
func (cq *CircularQueue) Rear() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	return cq.items[cq.rear], true
}

// IsEmpty returns true if the queue is empty.
// Time Complexity: O(1)
func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}

// IsFull returns true if the queue is at capacity.
// Time Complexity: O(1)
func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

// Size returns the current number of elements.
// Time Complexity: O(1)
func (cq *CircularQueue) Size() int {
	return cq.size
}

// ----------------------------------------------------------------------------
// Deque - Double-ended queue
// ----------------------------------------------------------------------------

// Deque implements a double-ended queue where elements can be
// added or removed from both ends.
type Deque struct {
	items []int
}

// NewDeque creates and returns a new empty deque.
func NewDeque() *Deque {
	return &Deque{items: []int{}}
}

// PushFront adds an element to the front of the deque.
// Time Complexity: O(n)
func (d *Deque) PushFront(val int) {
	d.items = append([]int{val}, d.items...)
}

// PushBack adds an element to the back of the deque.
// Time Complexity: O(1) amortized
func (d *Deque) PushBack(val int) {
	d.items = append(d.items, val)
}

// PopFront removes and returns the front element.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(n)
func (d *Deque) PopFront() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front, true
}

// PopBack removes and returns the back element.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(1)
func (d *Deque) PopBack() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	back := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return back, true
}

// PeekFront returns the front element without removing it.
// Time Complexity: O(1)
func (d *Deque) PeekFront() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	return d.items[0], true
}

// PeekBack returns the back element without removing it.
// Time Complexity: O(1)
func (d *Deque) PeekBack() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	return d.items[len(d.items)-1], true
}

// IsEmpty returns true if the deque is empty.
// Time Complexity: O(1)
func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

// Size returns the number of elements.
// Time Complexity: O(1)
func (d *Deque) Size() int {
	return len(d.items)
}

// ----------------------------------------------------------------------------
// PriorityQueue - Heap-based priority queue (min-heap by default)
// ----------------------------------------------------------------------------

// PriorityQueue implements a min-heap based priority queue.
// The smallest element has the highest priority.
type PriorityQueue struct {
	items []int
}

// NewPriorityQueue creates and returns a new empty priority queue.
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{items: []int{}}
}

// Push adds an element to the priority queue.
// Time Complexity: O(log n)
func (pq *PriorityQueue) Push(val int) {
	pq.items = append(pq.items, val)
	pq.siftUp(len(pq.items) - 1)
}

// Pop removes and returns the highest priority (smallest) element.
// Returns the value and true if successful, or 0 and false if empty.
// Time Complexity: O(log n)
func (pq *PriorityQueue) Pop() (int, bool) {
	if len(pq.items) == 0 {
		return 0, false
	}

	minVal := pq.items[0]
	lastIdx := len(pq.items) - 1
	pq.items[0] = pq.items[lastIdx]
	pq.items = pq.items[:lastIdx]

	if len(pq.items) > 0 {
		pq.siftDown(0)
	}

	return minVal, true
}

// Peek returns the highest priority element without removing it.
// Time Complexity: O(1)
func (pq *PriorityQueue) Peek() (int, bool) {
	if len(pq.items) == 0 {
		return 0, false
	}
	return pq.items[0], true
}

// IsEmpty returns true if the priority queue is empty.
// Time Complexity: O(1)
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.items) == 0
}

// Size returns the number of elements.
// Time Complexity: O(1)
func (pq *PriorityQueue) Size() int {
	return len(pq.items)
}

// siftUp moves an element up the heap to maintain heap property.
func (pq *PriorityQueue) siftUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if pq.items[idx] >= pq.items[parent] {
			break
		}
		pq.items[idx], pq.items[parent] = pq.items[parent], pq.items[idx]
		idx = parent
	}
}

// siftDown moves an element down the heap to maintain heap property.
func (pq *PriorityQueue) siftDown(idx int) {
	n := len(pq.items)
	for {
		smallest := idx
		left := 2*idx + 1
		right := 2*idx + 2

		if left < n && pq.items[left] < pq.items[smallest] {
			smallest = left
		}
		if right < n && pq.items[right] < pq.items[smallest] {
			smallest = right
		}

		if smallest == idx {
			break
		}

		pq.items[idx], pq.items[smallest] = pq.items[smallest], pq.items[idx]
		idx = smallest
	}
}

// ----------------------------------------------------------------------------
// MaxPriorityQueue - Max-heap based priority queue
// ----------------------------------------------------------------------------

// MaxPriorityQueue implements a max-heap based priority queue.
// The largest element has the highest priority.
type MaxPriorityQueue struct {
	items []int
}

// NewMaxPriorityQueue creates and returns a new empty max priority queue.
func NewMaxPriorityQueue() *MaxPriorityQueue {
	return &MaxPriorityQueue{items: []int{}}
}

// Push adds an element to the priority queue.
// Time Complexity: O(log n)
func (mpq *MaxPriorityQueue) Push(val int) {
	mpq.items = append(mpq.items, val)
	mpq.siftUp(len(mpq.items) - 1)
}

// Pop removes and returns the highest priority (largest) element.
// Time Complexity: O(log n)
func (mpq *MaxPriorityQueue) Pop() (int, bool) {
	if len(mpq.items) == 0 {
		return 0, false
	}

	maxVal := mpq.items[0]
	lastIdx := len(mpq.items) - 1
	mpq.items[0] = mpq.items[lastIdx]
	mpq.items = mpq.items[:lastIdx]

	if len(mpq.items) > 0 {
		mpq.siftDown(0)
	}

	return maxVal, true
}

// Peek returns the highest priority element without removing it.
// Time Complexity: O(1)
func (mpq *MaxPriorityQueue) Peek() (int, bool) {
	if len(mpq.items) == 0 {
		return 0, false
	}
	return mpq.items[0], true
}

// IsEmpty returns true if the priority queue is empty.
func (mpq *MaxPriorityQueue) IsEmpty() bool {
	return len(mpq.items) == 0
}

// Size returns the number of elements.
func (mpq *MaxPriorityQueue) Size() int {
	return len(mpq.items)
}

func (mpq *MaxPriorityQueue) siftUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if mpq.items[idx] <= mpq.items[parent] {
			break
		}
		mpq.items[idx], mpq.items[parent] = mpq.items[parent], mpq.items[idx]
		idx = parent
	}
}

func (mpq *MaxPriorityQueue) siftDown(idx int) {
	n := len(mpq.items)
	for {
		largest := idx
		left := 2*idx + 1
		right := 2*idx + 2

		if left < n && mpq.items[left] > mpq.items[largest] {
			largest = left
		}
		if right < n && mpq.items[right] > mpq.items[largest] {
			largest = right
		}

		if largest == idx {
			break
		}

		mpq.items[idx], mpq.items[largest] = mpq.items[largest], mpq.items[idx]
		idx = largest
	}
}
