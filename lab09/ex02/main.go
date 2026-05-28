/*
2. Чтение и сумма: Прочитайте числа из файла "numbers.txt", вычислите их
сумму и выведите на экран.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Println("Сумма чисел из файла:", sum)
}
