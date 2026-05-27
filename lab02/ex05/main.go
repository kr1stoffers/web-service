/*
5. Замыкание-аккумулятор: Создайте функцию adder, которая принимает
начальное значение и возвращает функцию, прибавляющую свой аргумент к
этому значению.
*/
package main

import "fmt"

func adder(initialValue int) func(int) int {
	sum := initialValue
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	acc := adder(10)
	fmt.Println(acc(5))
	fmt.Println(acc(3))
	fmt.Println(acc(10))
}
