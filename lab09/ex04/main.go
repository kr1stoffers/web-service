/*
4. Подсчёт строк: Напишите функцию, которая принимает имя файла и
возвращает количество строк в нём. Используйте bufio.Scanner.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesCount := 0
	for scanner.Scan() {
		linesCount++
	}
	return linesCount, scanner.Err()
}

func main() {
	count, err := countLines("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Количество строк в файле:", count)
}
