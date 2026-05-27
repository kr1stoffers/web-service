/*
2. Проверка на чётность: Создайте функцию isEven, возвращающую true, если
число чётное.
*/
package main

import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println(isEven(4))
	fmt.Println(isEven(7))
}
