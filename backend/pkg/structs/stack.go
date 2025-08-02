package structs

import (
	"fmt"
	"slices"
)

type Stack[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Push(T)
	Pop() T
	Print()
}

type stack[T any] struct {
	items []T
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{items: make([]T, 0)}
}

func (s *stack[T]) Size() int {
	return len(s.items)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *stack[T]) Clear() {
	s.items = make([]T, 0)
}

func (s *stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *stack[T]) Pop() T {
	if s.Size() == 0 {
		panic("cannot Pop from empty stack")
	}

	value := s.items[s.Size()-1]
	s.items = slices.Delete(s.items, s.Size()-1, s.Size())
	return value
}

func (s *stack[T]) Print() {
	fmt.Println(s.items)
}
