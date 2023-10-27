package stack

// A backwards queue
// A <- B <- C <- head
// push to head pop from head (LIFO)

type node struct {
	prev  *node
	value int
}

type Stack struct {
	head   *node
	length int
}

func (s *Stack) Push(value int) {
	node := node{value: value, prev: nil}
	if s.head == nil {
		s.length = 1
		s.head = &node
		return
	}
	z := node
	c := s.head
	z.prev = c
	s.head = &z
	s.length++
}

func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}
	tmpHead := s.head
	s.head = s.head.prev
	s.length--
	return tmpHead.value
}
