package binarytree

import (
	"fmt"
	"testing"
)

// func TestInsert(t *testing.T) {
// 	tree := new(Tree)
// 	tree.Insert()
// }

// depth first
func printTree(n *node) {
	if n == nil {
		return
	}
	printTree(n.left)
	fmt.Println(n.value)
	printTree(n.right)
}

func TestInsert(t *testing.T) {
	tree := Tree{}
	vals := []int{5, 3, 8, 1, 4, 7, 9}
	//		   (5)
	//		(3)	  (8)
	//	  (1)(4) (7)(9)
	for _, v := range vals {
		tree.Insert(v)
	}
	printTree(tree.root)
}

func TestGetRoot(t *testing.T) {
	tree := Tree{}
	tree.Insert(69)
	val := tree.GetRoot()
	if val != 69 {
		t.Errorf("GetRoot(), Val at root should be 69, got %d\n", val)
	}
	fmt.Printf("GetRoot(). Val at root should be 69, got %d\n", val)
}
