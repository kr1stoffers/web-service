/*
7. Обработка ошибок: Напишите функцию safeDiv, которая возвращает
результат деления и ошибку, если делитель равен нулю. В main обработайте
ошибку.
*/
package main

import (
	"errors"
	"fmt"
)

func safeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}

func main() {
	res1, err1 := safeDiv(10, 0)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := safeDiv(10, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}
}
