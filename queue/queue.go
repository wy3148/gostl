package queue

import (
	"container/list"
	"errors"
)

//QueueIf queue interface
type QueueIf interface {
	Empty() bool
	Size() int
	Front() (interface{}, error)
	Back() (interface{}, error)
	Push(e interface{})
	Pop()
}

//Queue the Queue struct exported
type queue struct {
	l *list.List
}

//NewQueue call this func to create a new queue
func NewQueue() QueueIf {
	return &queue{
		l: list.New(),
	}
}

//Empty return true if queue has no element,otherwise return false
func (q *queue) Empty() bool {
	if q.l.Len() == 0 {
		return true
	}
	return false
}

//Size return current size of the queue
func (q *queue) Size() int {
	return q.l.Len()
}

//Front return the first element in the queue
func (q *queue) Front() (interface{}, error) {
	if q.l.Len() == 0 {
		return nil, errors.New("The queue is emtpy")
	}
	return q.l.Front().Value, nil
}

//Back return the last element in the queue
func (q *queue) Back() (interface{}, error) {
	if q.l.Len() == 0 {
		return nil, errors.New("The queue is emtpy")
	}
	return q.l.Back().Value, nil
}

//Push a new element value
func (q *queue) Push(e interface{}) {
	q.l.PushBack(e)
}

//Pop pop the element from the queue
func (q *queue) Pop() {
	q.l.Remove(q.l.Front())
}
