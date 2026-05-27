/*
10. Дополнительно: Изучите пакет reflect и напишите функцию, которая
выводит все методы интерфейса (или типа) с помощью рефлексии.
*/
package main

import (
	"fmt"
	"reflect"
)

type SimpleCalculator struct{}

func (SimpleCalculator) Add(a, b int) int      { return a + b }
func (SimpleCalculator) Subtract(a, b int) int { return a - b }

func printMethods(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Printf("Тип %s имеет методов: %d\n", t.Name(), t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" - Метод #%d: %s\n", i+1, method.Name)
	}
}

func main() {
	calc := SimpleCalculator{}
	printMethods(calc)
}
