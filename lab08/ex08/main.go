/*
8. Цепочка вызовов: Создайте несколько функций, каждая из которых может
вернуть ошибку. В main вызовите их и обработайте ошибки, завершая программу при
первой ошибке.
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

func step1() error { return nil }
func step2() error { return errors.New("сбой на шаге 2") }
func step3() error { return nil }

func main() {
	if err := step1(); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	if err := step2(); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	if err := step3(); err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	fmt.Println("Все шаги выполнены успешно!")
}
