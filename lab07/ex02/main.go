/*
2. Сумма элементов: Разбейте большой срез чисел на несколько частей и
запустите горутины для вычисления суммы каждой части. Соберите
результаты через канал и выведите общую сумму.
*/
package main

import "fmt"

func sumPart(numbers []int, ch chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	ch <- sum
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	parts := 3
	partSize := len(data) / parts
	ch := make(chan int)

	for i := 0; i < parts; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(data[start:end], ch)
	}

	totalSum := 0
	for i := 0; i < parts; i++ {
		totalSum += <-ch
	}

	fmt.Println("Общая сумма:", totalSum)
}
