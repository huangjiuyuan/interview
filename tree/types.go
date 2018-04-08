package tree

type Tree struct {
	left  *Tree
	right *Tree
	value int
}

func NewTree(v int) *Tree {
	t := &Tree{
		left:  nil,
		right: nil,
		value: v,
	}
	return t
}

func (t *Tree) Insert(v int) {
	if t == nil {
		t = &Tree{nil, nil, v}
	}

	if v < t.value {
		t.left.Insert(v)
	}
	t.right.Insert(v)
}

type BSTree struct {
	left  *BSTree
	right *BSTree
	value int
}

func NewBSTree(v int) *BSTree {
	t := &BSTree{
		left:  nil,
		right: nil,
		value: v,
	}
	return t
}

func (bst *BSTree) Insert(v int) {
	if bst == nil {
		bst = &BSTree{nil, nil, v}
	}

	if bst.value == v {
		return
	} else if bst.value < v {
		if bst.right == nil {
			bst.right = NewBSTree(v)
		} else {
			bst.right.Insert(v)
		}
	} else {
		if bst.left == nil {
			bst.left = NewBSTree(v)
		} else {
			bst.left.Insert(v)
		}
	}
}
