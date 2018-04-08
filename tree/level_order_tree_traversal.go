package tree

import (
	"fmt"

	"github.com/huangjiuyuan/interview/queue"
)

func PrintLevelOrder(t *Tree) {
	if t == nil {
		return
	}

	q := queue.New()
	q.Enqueue(t)
	for q.Len() != 0 {
		t = q.Dequeue().(*Tree)
		if t == nil {
			continue
		}
		fmt.Printf("%d ", t.value)

		if t.left != nil {
			q.Enqueue(t.left)
		}
		if t.right != nil {
			q.Enqueue(t.right)
		}
	}
	fmt.Println()
}
