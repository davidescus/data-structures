package list

import (
	"testing"
)

func TestNoValues(t *testing.T) {

}

func TestMoveNext(t *testing.T) {

}

func TestMovePrev(t *testing.T) {

}

func TestRemoveFront(t *testing.T) {

}

func TestRemoveMiddle(t *testing.T) {

}

func TestRemoveAll(t *testing.T) {

}

func TestRemoveAllAndRepopulate(t *testing.T) {

}

func TestRemoveBack(t *testing.T) {

}

func TestPushFront(t *testing.T) {

}

func TestPushBack(t *testing.T) {

	tt := []int{1, 2, 3, 4}
	l := add(tt)

	node := l.Front()
	count := 0
	for _, expect := range tt {
		if expect != node.Value {
			t.Errorf("Expect: %d but got: %d \n", expect, node.Value)
		}
		count++
		node = node.Next()
	}

	if len(tt) != count {
		t.Errorf("Expect: %d but got: %d \n", len(tt), count)
	}
}

func add(tt []int) *List {
	l := &List{}
	for _, v := range tt {
		l.PushBack(v)
	}
	return l
}

//func compare(list *List, holder []int) (error bool, expect, got interface{}) {
//	node := list.Front()
//	for _, value := range holder {
//		if value != node.Value {
//			return true, value, node.Value
//		}
//		node = node.Next()
//	}
//
//	return false, "", ""
//}
