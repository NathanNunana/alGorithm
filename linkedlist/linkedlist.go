package linkedlist

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

func (l *LinkedList) Insert(value interface{}) {
	newNode := node{data: value}
	if l.head == nil {
		l.head = &newNode
	}

}
