/*
8. Ошибки как интерфейс: Создайте собственную структуру ошибки,
реализующую интерфейс error. Используйте её в функции.
*/
package main

import "fmt"

type InvalidAgeError struct {
	Age int
}

func (e InvalidAgeError) Error() string {
	return fmt.Sprintf("недопустимый возраст: %d (должен быть от 0 до 120)", e.Age)
}

func checkAge(age int) error {
	if age < 0 || age > 120 {
		return InvalidAgeError{Age: age}
	}
	return nil
}

func main() {
	err := checkAge(-5)
	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
}
