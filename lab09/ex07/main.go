/*
7. Стандартный ввод-вывод: Напишите программу, которая читает из os.Stdin
и пишет в os.Stdout, но при этом заменяет все вхождения "foo" на "bar".
Используйте bufio.Scanner и fmt.Println.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Программа запущена. Пишите текст (для выхода нажмите Ctrl+C):")

	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := strings.ReplaceAll(line, "foo", "bar")
		fmt.Println(modifiedLine)
	}
}
