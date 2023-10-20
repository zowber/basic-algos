package queue

// head -> A -> B -> C -> tail
// push to tail pop from head (FIFO)

type node struct { // private
	next  *node
	value int
}

type Queue struct {
	head   *node
	tail   *node
	length int
}

func (q *Queue) Enqueue(value int) {
	node := node{value: value, next: nil}

	if q.tail == nil {
		q.head = &node
		q.tail = &node
		return
	}
	q.tail.next = &node
	q.tail = q.tail.next

	q.length = q.length + 1
}

func (q *Queue) Dequeue() (int, bool) {
	var val int
	if q.head == nil {
		q.tail = nil
		return val, false
	}

	val = q.head.value
	q.head = q.head.next

	q.length = q.length - 1

	return val, true

}

func (q Queue) Peek(position int) (int, bool) {
	var val int
	if q.head == nil {
		return val, false
	}

	val = q.head.value
	return val, true

}
