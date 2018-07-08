package priorityQueue

import (
	"errors"
	"strings"
)

// Similar as C++ priority queue, in interal, priority queue is implemented
// with binary heap
// 1) It’s a complete tree (All levels are completely filled except possibly the last level and the last level has all keys as left as possible).
//This property of Binary Heap makes them suitable to be stored in an array.
// 2) A Binary Heap is either Min Heap or Max Heap.
//In a Min Binary Heap, the key at root must be minimum among all keys present in Binary Heap. The same property must be recursively true for all nodes in Binary Tree. Max Binary Heap is similar to MinHeap.

const (
	defaultQueueInitSize = 11
)

//PriorityQueueIf queue interface
type PriorityQueueIf interface {
	Empty() bool
	Size() int
	Top() (interface{}, error)
	Push(e interface{})
	Pop()
}

//Compare function, priority compare
//Return: true, va1 >= va2 otherwise return false
type Compare func(va1 interface{}, va2 interface{}) bool

//heap, current ith element, its children are : 2i+1, and 2i +2, i starts from 0
type priorityQueue struct {
	elements []interface{}
	size     int
	cmp      Compare
}

func byteCmp(va1 interface{}, va2 interface{}) bool {
	return va1.(byte) >= va2.(byte)
}

func intCmp(va1 interface{}, va2 interface{}) bool {
	return va1.(int) >= va2.(int)
}

func int64Cmp(va1 interface{}, va2 interface{}) bool {
	return va1.(int64) >= va2.(int64)
}

func stringCmp(va1 interface{}, va2 interface{}) bool {
	return strings.Compare(va1.(string), va2.(string)) >= 0
}

//NewPriorityQueue create a new priority queue
//Can pass the Compare, if it's nil, will use default compare functions
//for known data types, supported types are int, int64, byte, string so far
//By default, it is max priority queue, the  element on the top has highest priority
func NewPriorityQueue(c Compare) PriorityQueueIf {
	return &priorityQueue{
		elements: make([]interface{}, defaultQueueInitSize),
		size:     0,
		cmp:      c,
	}
}

//Return true when queue is empty
func (q *priorityQueue) Empty() bool {
	return q.size == 0
}

//Return the size of current queue, number of elements in the queue
func (q *priorityQueue) Size() int {
	return q.size
}

//Return the top element
func (q *priorityQueue) Top() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("empty priority queue")
	}
	return q.elements[0], nil
}

//Push an element into the queue
func (q *priorityQueue) Push(e interface{}) {
	i := q.size
	if i > len(q.elements) {
		newElements := make([]interface{}, len(q.elements)*2)
		copy(newElements, q.elements)
		q.elements = newElements
	}
	q.siftup(i, e)
	q.size++
}

//Pop the first element in the queue
func (q *priorityQueue) Pop() {
	if q.size == 0 {
		return
	}
	q.elements[0] = q.elements[q.size-1]
	q.elements[q.size-1] = nil
	q.size--
	q.siftdown()
}

func (q *priorityQueue) siftdown() {
	if q.cmp != nil {
		q.siftdownUsingComparator()
	} else {
		q.siftdownWithDefaultComparator()
	}
}

func (q *priorityQueue) siftdownWithDefaultComparator() {
	if q.size == 0 {
		return
	}
	var cmp Compare
	switch q.elements[0].(type) {
	case byte:
		cmp = byteCmp
	case int:
		cmp = intCmp
	case int64:
		cmp = int64Cmp
	case string:
		cmp = stringCmp
	default:
		panic("Queue element doesn't have valid Comparator")
	}

	k := 0
	for k < q.size {
		l := k<<1 + 1

		if l >= q.size {
			break
		}

		if l < q.size && l+1 < q.size && cmp(q.elements[l+1], q.elements[l]) {
			l = l + 1
		}

		if cmp(q.elements[l], q.elements[k]) {
			q.elements[l], q.elements[k] = q.elements[k], q.elements[l]
			k = l
		} else {
			break
		}
	}
}

func (q *priorityQueue) siftdownUsingComparator() {
	if q.size == 0 {
		return
	}

	k := 0
	for k < q.size {
		l := k<<1 + 1

		if l >= q.size {
			break
		}

		//always swap with the larger child
		if l < q.size && l+1 < q.size && q.cmp(q.elements[l+1], q.elements[l]) {
			l = l + 1
		}

		if q.cmp(q.elements[l], q.elements[k]) {
			q.elements[l], q.elements[k] = q.elements[k], q.elements[l]
			k = l
		} else {
			break
		}
	}
}

func (q *priorityQueue) siftup(k int, e interface{}) {
	if q.cmp != nil {
		q.siftupUsingComparator(k, e)
	} else {
		q.siftupWithDefaultComparator(k, e)
	}
}

//k, the pos to be inserted, e is the new element to be inserted
func (q *priorityQueue) siftupWithDefaultComparator(k int, e interface{}) {
	var cmp Compare
	switch e.(type) {
	case byte:
		cmp = byteCmp
	case int:
		cmp = intCmp
	case int64:
		cmp = int64Cmp
	case string:
		cmp = stringCmp
	default:
		panic("Queue element doesn't have valid Comparator")
	}

	q.elements[k] = e
	pos := k
	for pos > 0 {
		parent := (pos - 1) / 2
		if cmp(q.elements[pos], q.elements[parent]) {
			q.elements[parent], q.elements[pos] = q.elements[pos], q.elements[parent]
			pos = parent
		} else {
			break
		}
	}
}

func (q *priorityQueue) siftupUsingComparator(k int, e interface{}) {
	q.elements[k] = e
	pos := k
	for pos > 0 {
		parent := (pos - 1) / 2
		if q.cmp(q.elements[pos], q.elements[parent]) {
			q.elements[parent], q.elements[pos] = q.elements[pos], q.elements[parent]
			pos = parent
		} else {
			break
		}
	}
}
