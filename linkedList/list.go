package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
	size int
}

func (l *List) PushBack(n *Node) {
	l.size++

	if l.head == nil {
		l.head = n
		return
	}

	back := l.Back()
	back.next = n
}

func (l *List) PushFront(n *Node) {
	l.size++

	if l.head == nil {
		l.head = n
		return
	}

	currentHead := l.head
	l.head = n
	l.head.next = currentHead
}

func (l *List) Back() *Node {
	n := l.head
	for n.next != nil {
		n = n.next
	}
	return n
}

func (l *List) Front() *Node {
	return l.head
}

func main() {

	var list List

	n := Node{nil, 10}
	n.next = nil
	n.value = 10
	list.PushFront(&n)

	n2 := new(Node)
	n2.next = nil
	n2.value = 20
	list.PushFront(n2)

	n3 := new(Node)
	n3.next = nil
	n3.value = 30
	list.PushFront(n3)

	n4 := new(Node)
	n4.next = nil
	n4.value = 40
	list.PushBack(n4)

	front := list.Front()

	for front != nil {
		fmt.Println("Reference: ", front)
		front = front.next
	}

	fmt.Println(list.size)

}
