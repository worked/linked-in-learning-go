package main

// DoublyLinkList sets a reference  to the previous and next nodes
// to each element, including the head and tail nodes
// head (current) -> tail (previous)
// tail (current) -> head (next)

import (
	"errors"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = true

type Node[T comparable] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// CreateNode appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size.
func (l *DoublyLinkedList[T]) CreateNode(index int, v T) error {
	currentSize := l.size // 0 at the start, as nodes are added, increment by 1
	if index > currentSize {
		return errors.New("index exceeds list size")
	}

	newElement := &Node[T]{
		value: v,
	}
	l.size = currentSize + 1 // manually manage size since no underlying array

	// list is empty
	if l.head == nil {
		l.head, l.tail = newElement, newElement // set head and tail to the newElement
		return nil
	}

	// change head
	if index == 0 {
		newElement.next = l.head
		l.head.prev, l.head = newElement, newElement
		return nil
	}

	// change tail
	if index == currentSize {
		newElement.prev = l.tail
		l.tail.next, l.tail = newElement, newElement
		return nil
	}

	// change middle - find element at index
	current := l.head
	for i := 1; i < index; i++ {
		current = current.next
	}
	newElement.prev = current
	newElement.next = current.next
	current.next.prev, current.next = newElement, newElement
	return nil
}

// AddNodes creates nodes in the list (aka in memory)
// Each node references the previous and next nodes
// Each node can be accessed by index or by value
func (l *DoublyLinkedList[T]) AddNodes(nodes []struct {
	index int
	value T
}) error {
	for _, e := range nodes {
		if err := l.CreateNode(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

// SetHead sets the value of the head node
func (l *DoublyLinkedList[T]) SetHead(v T) error {
	if l.head == nil {
		return errors.New("list is empty")
	}
	l.head.value = v
	return nil
}

// SetTail sets the value of the tail node
func (l *DoublyLinkedList[T]) SetTail(v T) error {
	if l.tail == nil {
		return errors.New("list is empty")
	}
	l.tail.value = v
	return nil
}

// GetNode returns a node by value or index
func (l *DoublyLinkedList[T]) GetNode(identifier interface{}) (*Node[T], error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	switch identifier.(type) {
	case int:
		return l.GetNodeByIndex(identifier.(int))
	case T:
		return l.GetNodeByValue(identifier.(T))
	}
	return nil, errors.New("invalid identifier type")
}

// GetNodeByIndex returns a node by index
func (l *DoublyLinkedList[T]) GetNodeByIndex(index int) (*Node[T], error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("index out of bounds")
	}
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current, nil
}

// GetNodeByValue returns a node by value
func (l *DoublyLinkedList[T]) GetNodeByValue(value T) (*Node[T], error) {
	current := l.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("value not found")
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> TAIL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "TAIL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func main() {
	testCases := []struct {
		index int
		value string
	}{
		{index: 0, value: "A"},
		{index: 1, value: "B"},
		{index: 2, value: "C"},
		{index: 3, value: "D"},
	}
	dl := &DoublyLinkedList[string]{}
	err := dl.AddNodes(testCases)
	//forwardPrint := dl.PrintForward()
	//reversePrint := dl.PrintReverse()

	err = dl.SetHead("G")
	err = dl.SetTail("H")

	// Get node by index
	node, err := dl.GetNode(0)
	if err != nil {
		fmt.Println("Error getting node by index:", err)
	} else {
		fmt.Println("Node at index 0:", node.value)
	}

	// Get node by value
	node, err = dl.GetNode("B")
	if err != nil {
		fmt.Println("Error getting node by value:", err)
	} else {
		fmt.Println("Node with value 'B':", node.value)
	}

	forwardPrint := dl.PrintForward()
	reversePrint := dl.PrintReverse()

	fmt.Println(forwardPrint)
	fmt.Println(reversePrint)
}
