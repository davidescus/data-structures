package list

/*
list package provide implementation for double liked list
*/

// Node ...
type Node struct {
	next  *Node
	prev  *Node
	Value interface{}
}

// Next move to next node if exists
func (n *Node) Next() *Node {
	return n.next
}

// Prev move to previous node
func (n *Node) Prev() *Node {
	return n.prev
}

// List is the basic structure that keep the current state
type List struct {
	head *Node
	back *Node
	size int
}

// Len is used to get current size of list
func (l *List) Len() int {
	return l.size
}

// Remove first, last, middle Node
func (l *List) Remove(n *Node) {

	if l.Len() < 1 || n == nil {
		return
	}

	// it is the first on list
	if n.prev == nil {
		if n.next != nil {
			n.next.prev = nil
		}
		l.head = n.next
		n.next = nil
		return
	}

	// it is the last on list
	if n.next == nil {
		n.prev.next = nil
		return
	}

	// on middle
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil

	l.size--
}

// Front reference for first Node of list
// nil if list is empty
func (l *List) Front() *Node {
	return l.head
}

// PushFront create new node in front of list
func (l *List) PushFront(v interface{}) *Node {
	n := &Node{nil, nil, v}
	l.size++
	if l.head != nil {
		n.next = l.head
		l.head.prev = n
	}
	l.head = n

	return n
}

// Back get the last node if exists
func (l *List) Back() *Node {
	return l.back
}

// PushBack add node at the end of list
func (l *List) PushBack(v interface{}) *Node {
	n := &Node{nil, nil, v}
	l.size++
	if l.head == nil {
		l.head = n
		l.back = n
		return n
	}
	n.prev = l.back
	l.back.next = n
	l.back = n

	return n
}
