/*
6. Проверка типа ошибки: Используя type assertion, определите, является ли
возвращённая ошибка вашим кастомным типом, и выведите дополнительные поля.
*/
package main

import "fmt"

type InvalidInputError struct {
	Input  string
	Reason string
}

func (e *InvalidInputError) Error() string {
	return e.Reason
}

func checkInput(s string) error {
	if s == "" {
		return &InvalidInputError{Input: s, Reason: "пустое значение"}
	}
	return nil
}

func main() {
	err := checkInput("")
	if err != nil {
		if inputErr, ok := err.(*InvalidInputError); ok {
			fmt.Printf("Обнаружена InvalidInputError!\nПоле Input: '%s'\nПоле Reason: %s\n", inputErr.Input, inputErr.Reason)
		} else {
			fmt.Println("Какая-то другая ошибка:", err)
		}
	}
}
