package doublylinkedlist

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
	l.length++
	if idx > l.length {
		return
	}
	if idx == l.length {
		l.Append(val)
		return
	}
	if idx == 0 {
		l.Prepend(val)
		return
	}

	z := node{value: val}
	cursor := l.head
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	// 1. point new node next to target
	// 2. point new node prev to target prev
	// 3. point prev node next to new node
	// 4. point target prev to new node
	z.next = cursor
	z.prev = cursor.prev
	if cursor.prev != nil {
		cursor.prev.next = &z
	}
	cursor.prev = &z
}

func (l *List) RemoveAt(idx int) {
	if idx > l.length {
		return
	}

	// if statements for removing at head or tail
	// and if list only has one node, etc

	cursor := l.head
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	// A <-> B <-> C
	// remove at B
	cursor.prev.next = cursor.next
	cursor.next.prev = cursor.prev
	l.length--
}

func (l *List) Prepend(val int) {
	l.length++
	z := node{value: val}
	if l.head == nil {
		l.head = &z
		l.tail = l.head
		return
	}
	z.next = l.head
	l.head.prev = &z
	l.head = &z
}

func (l *List) Append(val int) {
	l.length++
	z := node{value: val}
	if l.tail == nil {
		l.tail = &z
		l.head = &z
		return
	}
	z.prev = l.tail
	l.tail.next = &z
	l.tail = &z
}

func (l *List) GetAt(idx int) int {
	if idx > l.length {
		return -1
	}

	cursor := l.head
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	return cursor.value
}
