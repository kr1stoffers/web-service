/*
1. Максимум из трёх: Напишите функцию maxOfThree, принимающую три
целых числа и возвращающую максимальное.
*/
package main

import "fmt"

func maxOfThree(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func main() {
	fmt.Println(maxOfThree(5, 12, 8))
}
