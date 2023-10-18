package doubly_linked_list

// doubly linked list
type node struct {
	value int
	next  *node
	prev  *node
}

type List struct {
	head   *node
	length int
}

func (l *List) InsertAt(idx int, val int) {

}
