/*
2. Чтение файла: Напишите функцию, которая читает файл и возвращает его
содержимое в виде строки. Если файл не существует, верните ошибку. Используйте
os.ReadFile и обработайте ошибку.
*/
package main

import (
	"fmt"
	"os"
)

func readFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	content, err := readFileToString("non_existent_file.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println("Содержимое:", content)
}
