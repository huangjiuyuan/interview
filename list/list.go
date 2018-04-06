package list

import "fmt"

type List struct {
	head *node
	tail *node
}

type node struct {
	prev *node
	next *node
	key  interface{}
}

func (l *List) Insert(key interface{}) {
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	n := &node{
		prev: curr,
		next: nil,
		key:  key,
	}
	curr.next = n
	l.tail = n
}

func (l *List) Print() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v -> ", list.key)
		list = list.next
	}
	fmt.Println("end")
}

func (l *List) Reverse() {
	curr := l.head
	var prev *node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
