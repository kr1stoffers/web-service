/*
7. Квадраты чисел: Напишите функцию, которая принимает срез и возвращает
новый срез с квадратами чисел, но при этом изменяет исходный срез через
указатель (например, func squareSlice(slice *[]int)).
*/
package main

import "fmt"

func squareSlice(slice *[]int) []int {
	newSlice := make([]int, len(*slice))
	for i, val := range *slice {
		newSlice[i] = val * val
		(*slice)[i] = val * val
	}
	return newSlice
}

func main() {
	origin := []int{2, 3, 4}
	fmt.Println("Старый исходный:", origin)
	squared := squareSlice(&origin)
	fmt.Println("Новый:", squared)
	fmt.Println("Исходный:", origin)
}
