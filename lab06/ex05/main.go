/*
5. Проверка типа: Напишите функцию, которая принимает interface{} и
определяет, является ли он числом (int, float64 и т.д.).
*/
package main

import "fmt"

func isNumber(i interface{}) bool {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(isNumber(42))
	fmt.Println(isNumber(3.14))
	fmt.Println(isNumber("строка"))
}
