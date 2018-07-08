package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	v := []int{1, 2, 3}

	for _, e := range v {
		q.Push(e)
	}

	for i := 0; i < len(v); i++ {
		e, err := q.Front()
		q.Pop()

		if err != nil {
			t.Fatal(err)
		}

		if v[i] != e.(int) {
			t.Fatalf("Element value is not equal to the original value %d:%d", v[i], e.(int))
		}
	}
}

func TestQueueBack(t *testing.T) {
	q := NewQueue()
	v := []int{1, 2, 3}
	for _, e := range v {
		q.Push(e)
	}
	e, err := q.Back()

	if err != nil {
		t.Fatal(err)
	}

	if e.(int) != 3 {
		t.Fatal("The last element in the queue is wrong")
	}
}

func TestQueueComplexElement(t *testing.T) {
	type MyE struct {
		Age   int
		House []string
	}

	q := NewQueue()

	v := &MyE{
		Age:   28,
		House: []string{"h1", "h2", "h3"},
	}

	v2 := &MyE{
		Age:   29,
		House: []string{"h11", "h22"},
	}

	q.Push(v)
	q.Push(v2)
	q.Pop()

	//verify second element
	e, err := q.Front()
	if err != nil {
		t.Fatal("Failed to get the first element in the queue")
	}

	for k, h := range e.(*MyE).House {
		if h != v2.House[k] {
			t.Fatalf("The value is not equal to original element:%s:%s", h, v2.House[k])
		}
	}
}

func BenchmarkQueue(b *testing.B) {
	q := NewQueue()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}
