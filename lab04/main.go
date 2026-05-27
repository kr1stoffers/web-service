/*
Задания 2, 3, 4, 7
*/
package main

import (
	"fmt"
	"lab4/numutil"
	"lab4/strutil"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	// Задание 2. strutil
	srcStr := "Привет, Go!"
	fmt.Println("Исходная строка:", srcStr)
	fmt.Println("Переворот:", strutil.Reverse(srcStr))
	fmt.Println("Верхний регистр:", strutil.ToUpper(srcStr))

	// Задание 3. numutil
	num := 17
	fmt.Printf("Число %d простое? %t\n", num, numutil.IsPrime(num))
	fmt.Printf("Факториал 5: %d\n", numutil.Factorial(5))

	// Задание 4. Генерация UUID через внешнюю зависимость
	id := uuid.New()
	fmt.Println("Сгенерированный UUID:", id.String())

	// Задание 7. Использование пакета color
	color.Cyan("бирюзовый цвет")
	color.Red("красный цвет")
}
