package stack

import "errors"

//Nueva version de stack
type Stack[T any] struct {
	items []T
}

func NewStack[T any](longitud int) *Stack[T] {
	stacks := new(Stack[T])
	stacks.items = make([]T, 0, longitud)
	return stacks
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (v T, err error) {
	if s.IsEmpty() {
		return v, errors.New("La pila esta vacia")
	}

	output := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return output, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len((*s).items) == 0
}

func (s *Stack[T]) Front() (T, error) {
	output := s.items[0]
	return output, nil
}
