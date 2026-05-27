/*
9.  Генератор последовательности: Создайте замыкание, которое генерирует
последовательные числа, начиная с 1, при каждом вызове.
*/
package main

import "fmt"

func newGenerator() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	gen := newGenerator()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
}
