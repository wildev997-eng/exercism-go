package linkedlist

import (
	"errors"
)

// Define List and Node types here.
type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.tail = newNode
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode

	}
}

func (l *List) Push(v any) {
	newNode := &Node{Value: v}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode

	}
}

func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("There's no head to shift to")
	}

	value := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return value, nil
}

func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, errors.New("There's no tail to pop")
	}
	value := l.tail.Value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return value, nil
}

func (l *List) Reverse() {
	current := l.head

	for current != nil {
		current.next, current.prev = current.prev, current.next
		current = current.prev
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Count() int {
	count := 0
	for n := l.head; n != nil; n = n.next {
		count++
	}
	return count
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (l *List) Delete(v any) bool {
	for n := l.head; n != nil; n = n.next {
		if n.Value != v {
			continue
		}

		if n.prev == nil {
			l.head = n.next // removing the first node
		} else {
			n.prev.next = n.next
		}

		if n.next == nil {
			l.tail = n.prev // removing the last node
		} else {
			n.next.prev = n.prev
		}

		return true
	}
	return false
}
