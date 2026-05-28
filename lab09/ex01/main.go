/*
1. Запись чисел: Напишите программу, которая запрашивает у пользователя 5
целых чисел и сохраняет их в файл "numbers.txt" (по одному числу на
строку).
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	fmt.Println("Введите 5 целых чисел:")
	for i := 0; i < 5; i++ {
		var num int
		fmt.Printf("Число %d: ", i+1)
		fmt.Scanln(&num)
		fmt.Fprintf(file, "%d\n", num)
	}
	fmt.Println("Числа успешно записаны в numbers.txt")
}
