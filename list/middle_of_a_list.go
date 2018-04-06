package list

import "fmt"

func PrintMiddle(l *List) {
	fast := l.head
	slow := l.head

	if l.head != nil {
		for fast.next != nil {
			fast = fast.next.next
			slow = slow.next
		}
		fmt.Printf("The middle is %+v\n", slow.key)
	}
}
