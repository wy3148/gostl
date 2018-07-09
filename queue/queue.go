package queue

import (
	"container/list"
	"errors"
)

//Queue the Queue struct exported
type Queue struct {
	l *list.List
}

//NewQueue call this func to create a new queue
func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

//Empty return true if queue has no element,otherwise return false
func (q *Queue) Empty() bool {
	if q.l.Len() == 0 {
		return true
	}
	return false
}

//Size return current size of the queue
func (q *Queue) Size() int {
	return q.l.Len()
}

//Front return the first element in the queue
func (q *Queue) Front() (interface{}, error) {
	if q.l.Len() == 0 {
		return nil, errors.New("The queue is emtpy")
	}
	return q.l.Front().Value, nil
}

//Back return the last element in the queue
func (q *Queue) Back() (interface{}, error) {
	if q.l.Len() == 0 {
		return nil, errors.New("The queue is emtpy")
	}
	return q.l.Back().Value, nil
}

//Push a new element value
func (q *Queue) Push(e interface{}) {
	q.l.PushBack(e)
}

//Pop pop the element from the queue
func (q *Queue) Pop() {
	q.l.Remove(q.l.Front())
}
