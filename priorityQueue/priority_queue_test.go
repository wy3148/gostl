package priorityQueue

import (
	"testing"
)

func TestPriorityQueueBasic(t *testing.T) {
	q := NewPriorityQueue(nil)
	e := []int{100, 90, 200, 2, 300, 301}
	expected := []int{301, 300, 200, 100, 90, 2}

	for _, v := range e {
		q.Push(v)
	}

	if q.Empty() {
		t.Fatal("Queue shouldn't be empty")
	}

	if q.Size() != len(e) {
		t.Fatal("Queue size is not correct")
	}

	index := 0
	for !q.Empty() {
		val, err := q.Top()
		q.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if val.(int) != expected[index] {
			t.Fatalf("Element is not the expected value %d,current value:%d", expected[index],
				val.(int))
		}
		index++
	}
}

type TestData struct {
	d1 string
	d2 int
	d3 int64
}

func UseTestDataCompare(v1 interface{}, v2 interface{}) bool {
	s1 := v1.(*TestData)
	s2 := v2.(*TestData)
	return s1.d3 >= s2.d3
}

func TestGeneralQueue(t *testing.T) {
	d := []*TestData{
		&TestData{
			d1: "n1",
			d2: 1,
			d3: 65,
		},
		&TestData{
			d1: "n2",
			d2: 2,
			d3: 64,
		},
		&TestData{
			d1: "n3",
			d2: 3,
			d3: 68,
		},
	}

	expected := []int64{68, 65, 64}
	q := NewPriorityQueue(UseTestDataCompare)
	for _, e := range d {
		t.Log(e.d3)
		q.Push(e)
	}

	index := 0
	for !q.Empty() {
		d, err := q.Top()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(d.(*TestData).d3)

		if d.(*TestData).d3 != expected[index] {
			t.Fatal("Element value is wrong,expected:", expected[index])
		}
		q.Pop()
		index++
	}
}

func TestQueueString(t *testing.T) {
	data := []string{
		"aa", "bb", "cc",
	}

	expected := []string{
		"cc", "bb", "aa",
	}

	q := NewPriorityQueue(nil)
	for _, e := range data {
		q.Push(e)
	}

	index := 0
	for !q.Empty() {
		d, err := q.Top()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(d.(string))
		if d.(string) != expected[index] {
			t.Fatal("Element value is wrong,expected", expected[index])
		}
		q.Pop()
		index++
	}
}
