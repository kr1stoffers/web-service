package main

// Среднее арифметическое: Запросите у пользователя 5 чисел (можно
// использовать массив), вычислите их среднее арифметическое и выведите результат с
// двумя знаками после запятой.

import "fmt"

func main() {
	var numbers [5]float64
	var sum float64

	fmt.Println("Введите 5 чисел:")
	for i := range 5 {
		fmt.Printf("Число %d: ", i+1)
		fmt.Scanln(&numbers[i])
		sum += numbers[i]
	}

	mean := sum / 5.0
	fmt.Printf("Среднее арифметическое: %.2f\n", mean)
}
