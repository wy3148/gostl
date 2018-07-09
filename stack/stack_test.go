package stack

import (
	"testing"
)

func TestStackBasic(t *testing.T) {
	s := NewStack()

	data := []interface{}{nil, nil, nil}
	expected := []interface{}{nil, nil, nil}

	for _, d := range data {
		s.Push(d)
	}

	if s.Size() != len(data) {
		t.Fatal("The stack size is wrong, expected:", s.Size())
	}

	index := 0
	for !s.Empty() {
		if e, err := s.Top(); err != nil {
			t.Fatal(err)
		} else {
			s.Pop()
			if e != expected[index] {
				t.Fatal("The element value is wrong,expected:", expected[index])
			}
		}
		index++
	}

	data = []interface{}{"ian1", nil, "wang1"}
	expected = []interface{}{"wang1", nil, "ian1"}

	for _, e := range data {
		s.Push(e)
	}

	index = 0
	for !s.Empty() {
		e, err := s.Top()
		s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if e != nil {
			if e.(string) != expected[index].(string) {
				t.Fatal("Element value is wrong")
			}
		}
		index++
	}
}

func TestStackPointer(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{
		Name: "ian",
		Age:  20,
	}

	p2 := &Person{
		Name: "wang",
		Age:  30,
	}

	s := NewStack()

	s.Push(p1)
	s.Push(p2)

	p, err := s.Top()

	if err != nil {
		t.Fatal(err)
	}

	s.Pop()

	if p.(*Person).Name != p2.Name {
		t.Fatal("The element value is wrong,expected:", p2.Name)
	}

	p, err = s.Top()

	if err != nil {
		t.Fatal(err)
	}

	s.Pop()

	if p.(*Person).Name != p1.Name {
		t.Fatal("The element value is wrong,expected:", p2.Name)
	}
}
