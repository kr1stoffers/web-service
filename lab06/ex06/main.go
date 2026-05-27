/*
6. Калькулятор с интерфейсом: Определите интерфейс Operation с методом
Apply(a, b int) int. Реализуйте структуры Add, Subtract, Multiply.
*/
package main

import "fmt"

type Operation interface {
	Apply(a, b int) int
}

type Add struct{}

func (Add) Apply(a, b int) int { return a + b }

type Subtract struct{}

func (Subtract) Apply(a, b int) int { return a - b }

type Multiply struct{}

func (Multiply) Apply(a, b int) int { return a * b }

func main() {
	var op Operation

	op = Add{}
	fmt.Println("10 + 5 =", op.Apply(10, 5))

	op = Subtract{}
	fmt.Println("10 - 5 =", op.Apply(10, 5))

	op = Multiply{}
	fmt.Println("10 * 5 =", op.Apply(10, 5))
}
