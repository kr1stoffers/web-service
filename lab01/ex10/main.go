package main

// Дополнительно: Изучите пакет math и используйте функцию math.Sqrt для
// вычисления квадратного корня из введённого числа (обработайте отрицательное значение).

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Введите число: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Ошибка: невозможно извлечь корень из отрицательного числа!")
	} else {
		res := math.Sqrt(num)
		fmt.Printf("Квадратный корень из %.2f равен %.4f\n", num, res)
	}
}
