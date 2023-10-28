package singlylinkedlist

import (
	"testing"
	"fmt"
)

func TestInsert(t *testing.T) {
	cases := []struct {	
		in int
		want int
	}{
		{69, 69},
	}
	list := List{}
	for _, c := range cases {
		for i := 0; i < c.in; i++ {
			list.Insert(i)
		}
		got := list.length
		if got != c.want {
			t.Errorf("%d Insert(n)s, list.length == %d, want %d\n", c.in, got, c.want)   
		}
		fmt.Printf("%d Insert(n)s, list.length == %d, want %d\n", c.in, got, c.want)
	}

}
func TestInsertAt(t *testing.T) {
	cases := []struct{
		in int
		want int
	}{
		{ 15, 420 },
	}
	list := List{}
	for i := 0; i < 20; i++ {
		list.Insert(i)
	}
	for _, c := range cases {
		list.InsertAt(c.in, 420)
		got := list.GetAt(c.in)
		if got != c.want {
			t.Errorf("InsertAt(%d, 420), GetAt(%d) == %d, want %d\n", c.in, c.in, got, c.want)
		}
		fmt.Printf("InsertAt(%d, 420), GetAt(%d) == %d, want %d\n", c.in, c.in, got, c.want)
	}

}

func TestGetAt(t *testing.T) {
	cases := []struct {
		in int
		want int
	}{
		{ 10, 10 },
	}
	list := List{}
	for i := 0; i < 20; i++ {
		list.Insert(i)
	}
	for _, c := range cases {
		got := list.GetAt(c.in)
		if got != c.want {
			t.Errorf("GetAt(%d) == %d, want %d\n", c.in, got, c.want)
		}
		fmt.Printf("GetAt(%d) == %d, want %d\n", c.in, got, c.want)
	}
}

func TestRemoveAt(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{15, 16},
	}
	list := List{}
	for i := 0; i < 20; i++ {
		list.Insert(i)
	}
	for _, c := range cases {
		list.DeleteAt(c.in)
		got := list.GetAt(c.in)
		if got != c.want {
			t.Errorf("RemoveAt(%d), GetAt(%d) == %d, want %d", c.in, c.in, got, c.want)
		}
		fmt.Printf("RemoveAt(%d), GetAt(%d) == %d, want %d", c.in, c.in, got, c.want)
	}
}
