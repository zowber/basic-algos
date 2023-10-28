package singlylinkedlist

import "fmt"

type node struct {
	value int
	next  *node
}

type List struct {
	head   *node
	length int
}

func (l *List) Insert(value int) {
	node := node{}
	node.value = value
	l.length++
	if l.head == nil {
		l.head = &node
		return
	}
	var cursor = l.head
	for i := 0; i < l.length; i++ {
		if cursor.next == nil { // reached end of list
			cursor.next = &node
			return
		}
		cursor = cursor.next
	}
}

func (l *List) InsertAt(idx int, val int) {
	// node to insert (z)
	z := node{value: val}
	cursor := l.head
	var a, b *node
	for i := 0; i < l.length; i++ {
		if i == idx-1 {
			// get previous node (a)
			a = cursor
		}
		// get next node (b)
		if i == idx {
			b = cursor
			break
		}
		cursor = cursor.next
	}
	// update nexts
	// point new node to next node
	z.next = b
	// point previous node to new node
	a.next = &z
	l.length++
}

func (l *List) GetAt(idx int) int {
	cursor := l.head
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	return cursor.value
}

func (l *List) DeleteAt(idx int) {
	// walk to node to delete (b)
	cursor := l.head
	var a, b, c *node
	for i := 0; i < idx; i++ {
		// get prev node (a)
		if i == idx-2 {
			a = cursor
		}
		// get next node (c) (could be nil if b is at tail)
		if i == idx-1 {
			b = cursor
			c = cursor.next
		}
		cursor = cursor.next
	}
	// update nexts and destroy (b)
	// (a).next points to (c) or tail
	a.next = c
	// destroy (b)
	b.next = nil
	b = nil
	l.length--
}

func (l *List) Walk() {
	cursor := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}
