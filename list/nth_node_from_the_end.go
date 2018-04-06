package list

import "fmt"

func PrintNthFromEnd(l *List, n int) {
	fast := l.head
	slow := l.head

	if l.head != nil {
		for n > 0 {
			if fast == nil {
				fmt.Printf("End of the list reached\n")
			}
			fast = fast.next
			n--
		}

		for fast.next != nil {
			fast = fast.next.next
			slow = slow.next
		}
		fmt.Printf("The nth from the end is %+v\n", slow.key)
	}
}
