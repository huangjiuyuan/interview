package tree

import (
	"fmt"

	"github.com/huangjiuyuan/interview/queue"
)

func PrintDiagonal(t *Tree) {
	if t == nil {
		return
	}

	q := queue.New()
	q.Enqueue(t)
	for q.Len() != 0 {
		t = q.Dequeue().(*Tree)
		if t.left != nil {
			q.Enqueue(t.left)
		}
		fmt.Printf("%d ", t.value)
		for t.right != nil {
			t = t.right
			if t.left != nil {
				q.Enqueue(t.left)
			}
			fmt.Printf("%d ", t.value)
		}
	}
}
