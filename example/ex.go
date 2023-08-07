package main

import (
	"fmt"

	"github.com/Sigumaa/gostack"
)

func main() {
	s := gostack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	elem, _ := s.Pop()
	fmt.Println(elem)

	elem, _ = s.Peek()
	fmt.Println(elem)

	s.Clear()
	_, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}

	s.Clear()

	ss := gostack.New[string]()
	ss.Push("Hello")
	ss.Push("World")
	ss.Push("!")

	e, _ := ss.Pop()
	fmt.Println(e)

	e, _ = ss.Peek()
	fmt.Println(e)

	ss.Clear()
}
