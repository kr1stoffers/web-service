/*
10. Дополнительно: Реализуйте функцию, которая паникует при делении на ноль,
но восстанавливается и возвращает 0.
*/
package main

import "fmt"

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Поймали панику:", r)
			result = 0 // Перезаписываем именованный результат
		}
	}()

	// Деление целых чисел int на 0 в Go вызывает рантайм-панику автоматически
	result = a / b
	return result
}

func main() {
	res := safeDivide(10, 0)
	fmt.Println("Результат безопасного деления:", res)
}
