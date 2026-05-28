/*
3. Собственная ошибка: Создайте тип InvalidInputError с полями Input и
Reason. Используйте его в функции, проверяющей, что строка не пуста и не длиннее 10
символов.
*/
package main

import "fmt"

type InvalidInputError struct {
	Input  string
	Reason string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("ошибка ввода '%s': %s", e.Input, e.Reason)
}

func validateString(s string) error {
	if s == "" {
		return &InvalidInputError{Input: s, Reason: "строка не должна быть пустой"}
	}
	if len(s) > 10 {
		return &InvalidInputError{Input: s, Reason: "длина строки превышает 10 символов"}
	}
	return nil
}

func main() {
	err := validateString("оченьдлиннаястрока")
	if err != nil {
		fmt.Println(err)
	}
}
