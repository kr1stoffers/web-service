package main

// Числа Фибоначчи: Выведите первые 20 чисел Фибоначчи, используя цикл.

import "fmt"

func main() {
	n := 20
	f1, f2 := 0, 1

	fmt.Print("Первые 20 чисел Фибоначчи: ")
	for range n {
		fmt.Printf("%d ", f1)
		f1, f2 = f2, f1+f2
	}
	fmt.Println()
}
