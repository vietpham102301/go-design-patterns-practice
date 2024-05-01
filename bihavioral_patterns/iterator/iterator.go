package iterator

// Node represents a node in the linked list
type Node struct {
	value interface{}
	next  *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given value to the end of the list
func (ll *LinkedList) Append(value interface{}) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Iterator returns an iterator for the linked list
func (ll *LinkedList) Iterator() Iterator {
	return &LinkedListIterator{current: ll.head}
}

// Iterator interface defines methods for iterating over elements
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// LinkedListIterator is a concrete iterator for LinkedList
type LinkedListIterator struct {
	current *Node
}

// HasNext checks if there is a next node in the list
func (it *LinkedListIterator) HasNext() bool {
	return it.current != nil
}

// Next returns the value of the current node and advances the iterator
func (it *LinkedListIterator) Next() interface{} {
	if it.HasNext() {
		value := it.current.value
		it.current = it.current.next
		return value
	}
	return nil // Return nil if no next node
}
