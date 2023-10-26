package doublyLinkedList

import (
	"testing"
)

func TestDoublyLinkedListPrepend(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{15, 1},
	}
	for _, c := range cases {
		list := List{}
		list.Prepend(c.in)
		got := list.length
		if got != c.want {
			t.Errorf("Prepend(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDoublyLinkedListInsertAt(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{69, 69},
	}
	// create a list
	list := List{}
	// insert 100 nodes
	for i := 0; i < 99; i++ {
		list.Append(i)
	}
	for _, c := range cases {

	}
}
