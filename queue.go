package unlimchan

import "errors"

// Queue represents FIFO collection.
type Queue struct {
	size int
	head *queueItem
	tail *queueItem
}

type queueItem struct {
	val interface{}
	next *queueItem
}

// Size returns count of elements in queue.
func (q *Queue) Size() int {
	return q.size
}

// Enqueue adds value to the end of the Queue.
func (q *Queue) Enqueue(val interface{}) {
	q.size++

	if q.tail == nil {
		q.tail = &queueItem{val:val}
		q.head = q.tail
		return
	}

	q.tail.next = &queueItem{val:val}
	q.tail = q.tail.next
}

// Dequeue removes and returns value at the beginning of the Queue. Returns error if Queue is empty.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}

	q.size--

	ans := q.head.val
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return ans, nil
}