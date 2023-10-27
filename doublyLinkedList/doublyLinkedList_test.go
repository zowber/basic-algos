package doublyLinkedList

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedListPrepend(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{15, 15},
	}
	list := List{}
	// insert 100 nodes
	for i := 0; i < 99; i++ {
		list.Append(i)
	}
	for _, c := range cases {
		list.Prepend(c.in)
		got := list.head.value
		if got != c.want {
			t.Errorf("Prepend(%d) == Get(0), got %d, want %d", c.in, got, c.want)
		}
		fmt.Printf("Prepend(%d) == list.head.value, got %d, want %d\n", c.in, got, c.want)
	}
}

func TestDoublyLinkedListAppend(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{69, 69},
	}
	list := List{}
	// insert 100 nodes
	for i := 0; i < 99; i++ {
		list.Append(i)
	}
	for _, c := range cases {
		list.Append(c.in)
		got := list.tail.value
		if got != c.want {
			t.Errorf("Append(%d) == %d, want %d", c.in, got, c.want)
		}
		fmt.Printf("Append(%d) == %d, want %d\n", c.in, got, c.want)
	}
}

func TestDoublyLinkedListInsertAt(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{69, 420},
	}
	// create a list
	list := List{}
	// insert 100 nodes
	for i := 0; i < 99; i++ {
		list.Append(i)
	}
	for _, c := range cases {
		list.InsertAt(c.in, c.want)
		got := list.Get(69)
		if got != c.want {
			t.Errorf("InsertAt(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDoublyLinkedListGet(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{69, 69},
	}
	// create a list
	list := List{}
	// insert 100 nodes
	for i := 0; i < 100; i++ {
		list.Append(i)
	}
	for _, c := range cases {
		got := list.Get(69)
		if got != c.want {
			t.Errorf("Get(%d) == %d, want %d", c.in, got, c.want)
		}
		fmt.Printf("Get(%d) == %d, want %d\n", c.in, got, c.want)
	}
}
