package gostack_test

import (
	"testing"

	"github.com/Sigumaa/gostack"
)

func TestStack(t *testing.T) {
	s := gostack.New[int]()

	if s.Len() != 0 {
		t.Errorf("Expected length of 0, got %d", s.Len())
	}

	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_, err = s.Peek()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	if s.Len() != 5 {
		t.Errorf("Expected length of 5, got %d", s.Len())
	}

	elem, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 5 {
		t.Errorf("Expected popped element to be 5, got %d", elem)
	}

	if s.Len() != 4 {
		t.Errorf("Expected length of 4, got %d", s.Len())
	}

	elem, err = s.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 4 {
		t.Errorf("Expected peeked element to be 4, got %d", elem)
	}

	if s.Len() != 4 {
		t.Errorf("Expected length of 4, got %d", s.Len())
	}

	s.Clear()

	if s.Len() != 0 {
		t.Errorf("Expected length of 0, got %d", s.Len())
	}

}
