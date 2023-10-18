package singly_linked_list

import "fmt"

type node struct {
	value int
	next  *node
}

type List struct {
	head   *node
	length int
}

func (sll *List) Insert(value int) {

	node := node{}
	node.value = value

	if sll.head == nil {
		sll.head = &node
		sll.length++
		return
	}

	if sll.head != nil {
		var cursor = sll.head
		for i := 0; i < sll.length; i++ {
			if cursor.next == nil { // reached end of list
				cursor.next = &node
				sll.length++
				return
			}
			cursor = cursor.next
		}
	}

}

func (sll *List) InsertAt(idx int, val int) {

	// node to insert (z)
	var z node
	z.value = val

	cursor := sll.head
	var a *node
	var b *node
	for i := 0; i < sll.length; i++ {
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

	sll.length++
}

func (sll *List) DeleteSl() {
	// walk to node to delete (b)
	// get previous node (a)
	// get next node (c) (could be nil if b is at tail)
	// update nexts and destroy (b) ??will the go just GC b after it's removed from the list??
	// (a).next points to (c)
	//     or tail
	// set (b) to nil
}

func (sll *List) Walk() {
	cursor := sll.head
	for i := 0; i < sll.length; i++ {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}
