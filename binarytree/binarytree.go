package binarytree

type node struct {
	value int
	left  *node
	right *node
}

type Tree struct {
	root *node
}

func (t *Tree) Insert(val int) {
	if t.root == nil {
		t.root = &node{val, nil, nil}
	} else {
		t.root.insert(val)
	}
}

func (n *node) insert(val int) {
	if val <= n.value {
		// try to insert left
		if n.left == nil {
			n.left = &node{val, nil, nil}
		} else {
			// recurse
			n.left.insert(val)
		}

	} else {
		// insert to insert right
		if n.right == nil {
			n.right = &node{val, nil, nil}
		} else {
			// recurse
			n.right.insert(val)
		}
	}
}

func (t *Tree) GetRoot() int {
	if t.root != nil {
		return t.root.value
	}
	return -1
}
