/*
8. Аналог tee: Реализуйте программу, которая читает из os.Stdin и
одновременно пишет в файл и в os.Stdout (используйте io.MultiWriter или
несколько Write).
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("tee_output.txt")
	defer file.Close()

	// Объединяем вывод в консоль и в файл
	multiWriter := io.MultiWriter(os.Stdout, file)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите текст (он запишется в файл и выведется в консоль):")

	for scanner.Scan() {
		line := scanner.Text() + "\n"
		_, _ = io.WriteString(multiWriter, line)
	}
}
