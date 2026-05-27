/*
3. Сумма и произведение: Напишите функцию, которая принимает срез чисел и
возвращает сумму и произведение.
*/
package main

import "fmt"

func calcSumAndProd(numbers []int) (int, int) {
	sum := 0
	prod := 1
	if len(numbers) == 0 {
		return 0, 0
	}
	for _, val := range numbers {
		sum += val
		prod *= val
	}
	return sum, prod
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sum, prod := calcSumAndProd(slice)
	fmt.Println(sum, prod)
}
