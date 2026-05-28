/*
4. defer с несколькими вызовами: Напишите функцию, которая открывает два
файла, записывает в них данные и закрывает с помощью defer. Продемонстрируйте
порядок выполнения defer (LIFO).
*/
package main

import (
	"fmt"
	"os"
)

func demonstrateDefer() {
	fmt.Println("Открываем файл 1")
	f1, _ := os.Create("file1.txt")
	defer func() {
		f1.Close()
		fmt.Println("Закрыт файл 1 (вызов defer)")
	}()

	fmt.Println("Открываем файл 2")
	f2, _ := os.Create("file2.txt")
	defer func() {
		f2.Close()
		fmt.Println("Закрыт файл 2 (вызов defer)")
	}()

	fmt.Println("Выполняем запись...")
}

func main() {
	demonstrateDefer()
}
