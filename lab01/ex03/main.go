package main

// Калькулятор: Напишите программу, которая запрашивает два числа и
// операцию (+, -, *, /) и выводит результат. Используйте switch для выбора операции.
// Обработайте деление на ноль.

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число, операцию (+,-,*,/) и второе число через пробел: ")
	fmt.Scanln(&a, &op, &b)

	switch op {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}
