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

// Return reference for the next node
// nil if not exists (list is empty or it is the last node in list)
func (n *Node) Next() *Node {
	return n.next
}

// Return reference for the previous node
// nil if not exists (list is empty or it is the first node in list)
func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	back *Node
	size int
}

// Return int, number of elements existing on list
func (l *List) Len() int {
	return l.size
}

// Remove node form list
func (l *List) Remove(n *Node) {

	if l.Len() < 1 || n == nil {
		return
	}

	// it is the first on list
	if n.prev == nil {
		n.next.prev = nil
		next := n.next
		n.next = nil
		l.head = next
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

// Front reference for first node of list
// nil if list is empty
func (l *List) Front() *Node {
	return l.head
}

// Add new node on front of list
// Return reference for new added node
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

// Return reference for last node of list
// nil if list is empty
func (l *List) Back() *Node {
	return l.back
}

// Add new node at the end of list
// return reference for new added node
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
