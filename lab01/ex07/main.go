package main

// Простые числа: Напишите программу, которая выводит все простые числа от
// 2 до N (N вводится пользователем).

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scanln(&n)

	fmt.Printf("Простые числа от 2 до %d: ", n)
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
