package pkg

import "container/list"

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Dequeue() interface{} {
	front := q.list.Front()
	if front != nil {
		q.list.Remove(front)
		return front.Value
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	front := q.list.Front()
	if front == nil {
		return true
	} else {
		return false
	}
}
