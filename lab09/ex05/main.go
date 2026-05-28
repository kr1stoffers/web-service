/*
5. Поиск в файле: Прочитайте файл и выведите все строки, содержащие
заданную подстроку (ввод подстроки с консоли).
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var query string
	fmt.Print("Введите подстроку для поиска: ")
	fmt.Scanln(&query)

	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Найденные строки:")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, query) {
			fmt.Println(line)
		}
	}
}
