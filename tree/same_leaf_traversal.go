package tree

func LeafTraversalIsSame(t1, t2 *Tree) bool {
	if t1 == nil || t2 == nil {
		return false
	}

	var traverse func(t *Tree, s []int)
	traverse = func(t *Tree, s []int) {
		if t == nil {
			return
		}
		if t.IsLeaf() {
			s = append(s, t.value)
		}
		traverse(t.left, s)
		traverse(t.right, s)
	}

	s1 := make([]int, 10)
	s2 := make([]int, 10)
	traverse(t1, s1)
	traverse(t2, s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
