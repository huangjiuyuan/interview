package tree

import (
	"fmt"

	"github.com/huangjiuyuan/interview/queue"
)

func PrintLineByLine(t *Tree) {
	if t == nil {
		return
	}

	q1 := queue.New()
	q2 := queue.New()
	q1.Enqueue(t)
	for q1.Len() != 0 && q2.Len() != 0 {
		for q1.Len() != 0 {
			t = q1.Dequeue().(*Tree)
			fmt.Printf("%d ", t.value)
			if t.left != nil {
				q2.Enqueue(t.left)
			}
			if t.right != nil {
				q2.Enqueue(t.right)
			}
		}
		fmt.Println()

		for q2.Len() != 0 {
			t = q2.Dequeue().(*Tree)
			fmt.Printf("%d ", t.value)
			if t.left != nil {
				q1.Enqueue(t.left)
			}
			if t.right != nil {
				q1.Enqueue(t.right)
			}
		}
		fmt.Println()
	}
}
