/*
6. Стек на срезе: Реализуйте структуру Stack с методами Push, Pop, IsEmpty.
Используйте срез в качестве внутреннего хранилища.
*/
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(val int) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("стек пуст")
	}
	index := len(s.elements) - 1
	val := s.elements[index]
	s.elements = s.elements[:index]
	return val, nil
}

func (s Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)

	val, _ := s.Pop()
	fmt.Println(val)
	val, _ = s.Pop()
	fmt.Println(val)

	fmt.Println(s.IsEmpty())
}
