/*
1. Функция Sqrt: Напишите функцию Sqrt, которая возвращает ошибку при
отрицательном аргументе. Используйте errors.New.
*/
package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("нельзя извлекать квадратный корень из отрицательного числа")
	}
	return math.Sqrt(x), nil
}

func main() {
	r, e := Sqrt(2)
	if e != nil {
		fmt.Println("Ошибка:", e)
		return
	}
	fmt.Println(r)

	res, err := Sqrt(-4)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(res)

}
