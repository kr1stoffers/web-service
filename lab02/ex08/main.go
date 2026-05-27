/*
8. Функция высшего порядка: Напишите функцию mapFunc, которая применяет
переданную функцию к каждому элементу массива и возвращает новый
массив.
*/
package main

import "fmt"

func mapFunc(arr [5]int, f func(int) int) [5]int {
	var result [5]int
	for i, val := range arr {
		result[i] = f(val)
	}
	return result
}

func main() {
	input := [5]int{1, 2, 3, 4, 5}
	double := func(x int) int { return x * 2 }
	res := mapFunc(input, double)
	fmt.Println(res)
}
