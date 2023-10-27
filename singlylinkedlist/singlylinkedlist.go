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
	var a *node
	var b *node
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

func (l *List) DeleteAt(idx int) {
	// walk to node to delete (b)
	// get previous node (a)
	// get next node (c) (could be nil if b is at tail)
	// update nexts and destroy (b) ??will the go just GC b after it's removed from the list??
	// (a).next points to (c)
	//     or tail
	// set (b) to nil
}

func (l *List) Walk() {
	cursor := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}
