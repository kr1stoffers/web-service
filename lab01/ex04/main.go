package main

// Таблица умножения: Выведите таблицу умножения от 1 до 10 в виде матрицы 10×10.

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
