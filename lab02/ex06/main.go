/*
6. Рекурсивный обход: Напишите рекурсивную функцию для вычисления чисел
Фибоначчи (с мемоизацией или без).
*/
package main

import "fmt"

var memo = map[int]int{0: 0, 1: 1}

func fibonacci(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func main() {
	fmt.Println(fibonacci(40))
}
