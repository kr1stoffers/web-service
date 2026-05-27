/*
4. Напишите функцию isPrime целого типа, возвращающую 1, если целый
параметр N (N > 1) является простым числом, и 0 в противном случае.
*/
package main

import "fmt"

func isPrime(n int) int {
	if n <= 1 {
		return 0
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(15))
}
