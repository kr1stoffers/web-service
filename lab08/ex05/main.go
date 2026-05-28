/*
5. Паника и восстановление: Напишите функцию, которая вызывает панику при
определённом условии, но восстанавливается через recover и возвращает ошибку
вместо паники.
*/
package main

import (
	"fmt"
)

func processData(val int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("программа восстановилась после паники: %v", r)
		}
	}()

	if val < 0 {
		panic("критическое состояние: значение меньше нуля")
	}
	return nil
}

func main() {
	err := processData(-10)
	if err != nil {
		fmt.Println("Обработанная ошибка:", err)
	}
}
