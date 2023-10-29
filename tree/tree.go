package tree

type leaf struct {
	value    int
	children *leaf
}

type Tree struct {
	root *leaf
}

func (t *Tree) Insert() {

}
