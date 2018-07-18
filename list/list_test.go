package list

import (
	"testing"
)

func TestNoValues(t *testing.T) {
	name := "Test No Values"

	l := &List{}

	if l.Front() != nil || l.Back() != nil || l.Len() != 0 {
		t.Errorf("%s\nFor empty list must have len = 0 head, back = nil", name)
	}
}

func TestRemoveFront(t *testing.T) {
	name := "Test Remove Front"
	exp := []int{2, 3}

	l := &List{}
	n := l.PushBack(1)
	_ = l.PushBack(2)
	_ = l.PushBack(3)

	l.Remove(n)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestRemoveMiddle(t *testing.T) {
	name := "Test Remove Middle"
	exp := []int{1, 3}

	l := &List{}
	_ = l.PushBack(1)
	n2 := l.PushBack(2)
	_ = l.PushBack(3)

	l.Remove(n2)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestRemoveAllOneElement(t *testing.T) {
	name := "Test Remove All One Element"
	var exp []int

	l := &List{}
	n1 := l.PushBack(1)

	l.Remove(n1)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestRemoveAllOneManyElements(t *testing.T) {
	name := "Test Remove All And Repopulate"
	var exp []int

	l := &List{}
	n1 := l.PushBack(1)
	n2 := l.PushBack(2)
	n3 := l.PushBack(3)

	l.Remove(n1)
	l.Remove(n2)
	l.Remove(n3)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestRemoveAllAndRepopulate(t *testing.T) {
	name := "Test Remove All And Repopulate"
	exp := []int{4, 5, 6}

	l := &List{}
	n1 := l.PushBack(1)
	n2 := l.PushBack(2)
	n3 := l.PushBack(3)

	l.Remove(n1)
	l.Remove(n2)
	l.Remove(n3)

	l.PushBack(4)
	l.PushBack(5)
	l.PushBack(6)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestRemoveBack(t *testing.T) {
	name := "Test Remove Back"
	exp := []int{1, 2}

	l := &List{}
	_ = l.PushBack(1)
	_ = l.PushBack(2)
	n := l.PushBack(3)

	l.Remove(n)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestPushFront(t *testing.T) {
	name := "Test PushFront"
	exp := []int{3, 2, 1}

	l := &List{}
	_ = l.PushFront(1)
	_ = l.PushFront(2)
	_ = l.PushFront(3)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestPushBack(t *testing.T) {
	name := "Test PushBack"
	exp := []int{1, 2, 3}

	l := &List{}
	_ = l.PushBack(1)
	_ = l.PushBack(2)
	_ = l.PushBack(3)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func TestPushFrontBack(t *testing.T) {
	name := "Test PushFront - PushBack"
	exp := []int{4, 2, 1, 3}

	l := &List{}
	_ = l.PushBack(1)
	_ = l.PushFront(2)
	_ = l.PushBack(3)
	_ = l.PushFront(4)

	got, err := compare(l, exp)
	if err == true {
		t.Errorf("%s\nExp: %v\nGot: %v\n", name, exp, got)
	}
}

func compare(l *List, exp []int) ([]int, bool) {
	var got []int
	var err bool

	for node := l.Front(); node != nil; node = node.Next() {
		got = append(got, node.Value.(int))
	}

	if len(got) != len(exp) {
		return got, true
	}

	for i, v := range exp {
		if v != got[i] {
			err = true
		}
	}

	return got, err
}
