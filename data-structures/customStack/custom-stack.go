package main

import (
	"errors"
)

// Generic stack implementation
type Stack[T any] struct {
	elements []T //elements slice of any type
}

// Adds a new element to the top of the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Returns a bool to indicate if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Reads and removes an element from the stack
// or returns an error of stack is empty.
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("Cannot pop from empty stack")
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return &top, nil
}

func main() {

}
