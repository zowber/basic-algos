package doublyLinkedList

// doubly linked list
type node struct {
	value int
	next  *node
	prev  *node
}

type List struct {
	head   *node
	tail   *node
	length int
}

func (l *List) InsertAt(idx int, val int) {
	if idx > l.length {
		return
	} else if idx == l.length {
		l.Append(val)
		return
	} else if idx == 0 {
		l.Prepend(val)
		return
	}
	var a, b *node
	z := node{value: val}
	cursor := l.head
	for i := 0; i < idx; i++ {
		if i == idx-1 {
			a = cursor
		}
		if i == idx {
			b = cursor
			break
		}
		cursor = cursor.next
	}
	// point new node to target
	z.next = b
	// point new node to target prev
	z.prev = b.prev
	// point target to new node
	b.prev = &z
	// point target prev to new node
	a.next = &z
	// increment length
	l.length++
}

func (l *List) Remove(val int) {

}

func (l *List) RemoveAt(idx int) {

}

func (l *List) Prepend(val int) {
	// Z -> A
	// Z <- A
	// Head -> Z
	z := node{value: val}
	// if first item in list
	// tail == head
	if l.head == nil {
		l.head = &z
		l.tail = l.head
		return
	}
	// new node points to current head
	z.next = l.head
	// current first node prev points to new node
	l.head.prev = &z
	// head points to new node
	l.head = &z
	l.length++
}

func (l *List) Append(val int) {
	// Z -> B
	// Z <- B
	// Tail -> Z
	z := node{value: val}
	// if first item in list
	// tail == head
	if l.tail == nil {
		l.tail = &z
		l.head = l.tail
		return
	}
	z.prev = l.tail
	l.tail.prev = &z
	l.tail = &z
	l.length++
}

func (l *List) Get(idx int) int {
	return 0
}
