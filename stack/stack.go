package stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	l *list.List
}

//NewStack create a new stack, stack interal is implemented
//in golang/list pakcage which is double linked list
func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

//Empty(), return true if the stack is empty
func (s *Stack) Empty() bool {
	return s.l.Len() == 0
}

//Size return the size of the elements in the stack
func (s *Stack) Size() int {
	return s.l.Len()
}

//Top return the last element in the stack
func (s *Stack) Top() (interface{}, error) {
	e := s.l.Back()
	if e == nil {
		return nil, errors.New("The stack is empty")
	}
	return e.Value, nil
}

//Push push an element into the stack
func (s *Stack) Push(e interface{}) {
	s.l.PushBack(e)
}

//Pop pop the element from the stacl
func (s *Stack) Pop() {
	e := s.l.Back()
	if e != nil {
		s.l.Remove(s.l.Back())
	}
}
