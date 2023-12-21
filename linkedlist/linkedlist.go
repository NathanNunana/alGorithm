package linkedlist

import (
	"fmt"
)

// node struct
type node struct {
	data interface{}
	next *node
}

// linkedList struct
type LinkedList struct {
	head *node
	len  int
}

// returns the lenth of the linked list
func (l LinkedList) GetLength() int { return l.len }

// adds a node at the beginning of the list
func (l *LinkedList) Prepend(value interface{}) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
	} else {
		second := l.head
		l.head = newNode
		l.head.next = second
	}
	l.len++
}

// adds the node at the end of the list
func (l *LinkedList) Append(value interface{}) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.len++
}

// removes a node by value
func (l *LinkedList) RemoveNode(value interface{}) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.len--
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
	}
}

// prints the entire linked list
func (l LinkedList) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%v\n", current.data)
		current = current.next
	}
}
