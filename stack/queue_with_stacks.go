package stack

type Queue struct {
	s1, s2 *Stack
	length int
}

func (q *Queue) Dequeue() interface{} {
	if q.s2.Len() > 0 {
		q.length--
		return q.s2.Pop()
	}
	if q.s1.Len() > 0 {
		for q.s1.Len() > 0 {
			q.s2.Push(q.s1.Pop())
		}
		q.length--
		return q.s2.Pop()
	}
	return nil
}

func (q *Queue) Enqueue(value interface{}) {
	q.s1.Push(value)
	q.length++
}
