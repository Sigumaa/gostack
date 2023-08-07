package gostack

import (
	"fmt"
)

type Stack[T any] struct {
	elements []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	s.elements = append(s.elements, elem)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Len() == 0 {
		return zero[T](), fmt.Errorf("Stack is empty")
	}
	elem := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return elem, nil
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func zero[T any]() T {
	var zero T
	return zero
}
