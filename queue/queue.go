package queue

import "github.com/lephuocmy668/challengers/linkedlist"

// Queue ...
type Queue struct {
	list *linkedlist.LinkedList
}

// NewQueue return new Queue
func NewQueue() *Queue {
	return &Queue{
		list: linkedlist.NewLinkedList(),
	}
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Peek ...
func (q *Queue) Peek() int {
	return q.list.Head.Val
}

// Enqueue ...
func (q *Queue) Enqueue(val int) {
	q.list.Append(val)
}

// Dequeue ...
func (q *Queue) Dequeue() *int {
	if q.list.IsEmpty() {
		return nil
	}

	res := q.list.Head.Val
	q.list.DeleteHead()
	return &res
}
