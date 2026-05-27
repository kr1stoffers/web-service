/*
1. Обмен значениями: Напишите функцию swap, которая принимает два
указателя на целые числа и меняет их значения местами.
*/
package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x, y)
}
